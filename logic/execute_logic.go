package logic

import (
	"context"
	"errors"
	"fmt"
	"gin_mani_center/clients"
	"gin_mani_center/dal"
	"gin_mani_center/model"
	pb_mani "gin_mani_center/pb"
	"gin_mani_center/util"
	logx "github.com/amoghe/distillog"
	"sync"
)

const SysOperatorUser = "sys"

func ExecuteRule(ctx context.Context, rule *pb_mani.Rule) (taskId string, err error) {
	var ruleId string
	var user string

	if rule.User == "" {
		user = SysOperatorUser
	} else {
		user = rule.User
	}
	ruleId, err = getRuleId(ctx, rule, user)
	if err != nil {
		logx.Errorf("ExecuteRule getRuleId error:%v", err)
		return "", err
	}
	engine_rpc, err := clients.GetEngineClient()
	if err != nil {
		logx.Errorf("ExecuteRule GetEngineClient error:%v", err)
		return "", err
	}
	taskReq := &pb_mani.CreateTaskReq{
		Task: &pb_mani.Task{
			RuleId:       ruleId,
			Operator:     user,
			ExecuteState: pb_mani.ExecuteState_execute_wait,
			OutputState:  pb_mani.FileState_file_invalid,
		},
	}
	TaskResp, err := engine_rpc.CreateTask(ctx, taskReq)
	if err != nil {
		logx.Errorf("ExecuteRule CreateTask error:%v", err)
		return
	}
	taskId = TaskResp.TaskId

	var wg sync.WaitGroup
	wg.Add(1)
	util.GoParallel(ctx, func() {
		defer wg.Done()
		updateTaskReq := &pb_mani.UpdateTaskReq{
			Task: &pb_mani.Task{
				TaskId:       taskId,
				ExecuteState: pb_mani.ExecuteState_execute_running,
			},
		}
		_, err := engine_rpc.UpdateTask(ctx, updateTaskReq)
		if err != nil {
			logx.Errorf("ExecuteRule UpdateTask error:%v", err)
			return
		}
	})
	wg.Add(1)
	util.GoParallel(ctx, func() {
		defer wg.Done()
		message := &model.Message{
			Id:       util.GenUID(),
			TaskId:   taskId,
			RuleId:   ruleId,
			Operator: user,
			Extra:    fmt.Sprintf("%s-%s-%s-%d-execute-rule-task", user, taskId, ruleId, rule.RuleType),
		}
		if err := dal.SendMessage(message); err != nil {
			logx.Errorf("ExecuteRule SendMessage error:%v", err)
			return
		}
	})
	wg.Wait()
	return taskId, nil

}

func getRuleId(ctx context.Context, rule *pb_mani.Rule, user string) (string, error) {
	if rule.RuleId != "" {
		rules, err := dal.GetRuleByCondition(ctx,
			rule.RuleId,
			"",
			pb_mani.RuleType_unknown_rule_type,
			pb_mani.RuleState_rule_state_valid,
			pb_mani.ExecuteType_unknown_exec_type)
		if err != nil {
			logx.Errorf("getRuleByRuleId dal GetRuleByCondition error:%v", err)
			return "", err
		}
		if len(rules) == 0 {
			logx.Errorf("getRuleByRuleId no rule")
			return "", errors.New("no rule")
		}
		return rules[0].RuleId, nil
	}
	ruleId := util.GenUID()
	ruleModel := &model.RuleModel{
		RuleId:      ruleId,
		User:        user,
		RuleType:    int64(rule.RuleType),
		RuleState:   int64(rule.RuleState),
		ExecuteType: int64(rule.ExecuteType),
		DescText:    rule.RuleConfig.DescText,
		SourceName:  rule.RuleConfig.SourceName,
		SourceUrl:   rule.RuleConfig.SourceUrl,
		SourceState: int64(rule.RuleConfig.SourceState),
	}
	if err := dal.AddRule(ctx, ruleModel); err != nil {
		logx.Errorf("getRuleId dal AddRule error:%v", err)
		return "", err
	}
	return ruleId, nil
}
