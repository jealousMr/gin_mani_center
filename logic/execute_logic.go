package logic

import (
	"context"
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
	ruleId = util.GenUID()
	var wg sync.WaitGroup

	wg.Add(1)
	util.GoParallel(ctx, func() {
		defer wg.Done()
		err = createNewRule(ctx, rule, user, ruleId)
		if err != nil {
			logx.Errorf("ExecuteRule createNewRule error:%v", err)
		}
	})

	wg.Add(1)
	util.GoParallel(ctx, func() {
		defer wg.Done()
		engine_rpc, err := clients.GetEngineClient()
		if err != nil {
			logx.Errorf("ExecuteRule GetEngineClient error:%v", err)
			return
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
	})
	wg.Wait()

	if err != nil {
		logx.Errorf("ExecuteRule error:%v", err)
		return
	}

	message := &model.Message{
		Id:       util.GenUID(),
		TaskId:   taskId,
		RuleId:   ruleId,
		Operator: user,
		Extra:    fmt.Sprintf("%s-%s-%s-%d-execute-rule-task", user, taskId, ruleId, rule.RuleType),
	}
	if err := dal.SendMessage(message); err != nil {
		logx.Errorf("ExecuteRule SendMessage error:%v", err)
		return "", nil
	}
	return taskId, nil

}

func createNewRule(ctx context.Context, rule *pb_mani.Rule, user, ruleId string) (err error) {
	currentDesc := rule.RuleConfig.DescText
	var ruleModel *model.RuleModel

	if rule.RuleId != "" {
		ruleModel, err = dal.GetRuleById(ctx, rule.RuleId)
		if err != nil {
			logx.Errorf("dal GetRuleById error:%v", err)
			return err
		}
		ruleModel.RuleId = ruleId
		ruleModel.User = user
		ruleModel.DescText = currentDesc
		ruleModel.RuleType = int64(rule.RuleType)
	} else {
		ruleModel = &model.RuleModel{
			RuleId:      ruleId,
			User:        user,
			RuleType:    int64(rule.RuleType),
			RuleState:   int64(rule.RuleState),
			ExecuteType: int64(rule.ExecuteType),
			DescText:    currentDesc,
			SourceName:  rule.RuleConfig.SourceName,
			SourceUrl:   rule.RuleConfig.SourceUrl,
			SourceState: int64(rule.RuleConfig.SourceState),
		}
	}
	if err := dal.AddRule(ctx, ruleModel); err != nil {
		logx.Errorf("dal AddRule error:%v", err)
		return err
	}
	return nil
}
