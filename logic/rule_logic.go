package logic

import (
	"context"
	"gin_mani_center/dal"
	"gin_mani_center/model"
	pb_mani "gin_mani_center/pb"
	"gin_mani_center/util"
	logx "github.com/amoghe/distillog"
)

func AddRule(ctx context.Context, req *pb_mani.AddRuleReq) (err error) {
	rule := &model.RuleModel{
		RuleId:      util.GenUID(),
		User:        req.Rule.User,
		RuleType:    int64(req.Rule.RuleType),
		RuleState:   int64(req.Rule.RuleState),
		ExecuteType: int64(req.Rule.ExecuteType),
		DescText:    req.Rule.RuleConfig.DescText,
		SourceName:  req.Rule.RuleConfig.SourceName,
		SourceUrl:   req.Rule.RuleConfig.SourceUrl,
		SourceState: int64(req.Rule.RuleConfig.SourceState),
	}
	if err = dal.AddRule(ctx, rule); err != nil {
		logx.Errorf("error logic AddRule")
		return err
	}
	return nil
}

func GetRuleByRuleType(ctx context.Context, req *pb_mani.GetRuleByRuleTypeReq) (resp *pb_mani.GetRuleByRuleTypeResp, err error) {
	limit, offset := util.GetLimitAndOffset(req.PageNo, req.PageSize)
	rules, total, err := dal.GetRuleByRuleType(ctx, req.RuleType, int(limit), int(offset))
	if err != nil {
		logx.Errorf("error logic GetRuleByRuleType")
		return nil, err
	}
	ruleList := make([]*pb_mani.Rule, 0)
	for _, rule := range rules {
		ruleList = append(ruleList, &pb_mani.Rule{
			RuleId:      rule.RuleId,
			User:        rule.User,
			RuleType:    pb_mani.RuleType(rule.RuleType),
			RuleState:   pb_mani.RuleState(rule.RuleState),
			ExecuteType: pb_mani.ExecuteType(rule.ExecuteType),
			RuleConfig: &pb_mani.RuleConfig{
				DescText:    rule.DescText,
				SourceUrl:   rule.SourceUrl,
				SourceName:  rule.SourceName,
				SourceState: pb_mani.FileState(rule.SourceState),
			},
		})
	}
	return &pb_mani.GetRuleByRuleTypeResp{
		RuleList: ruleList,
		Page: &pb_mani.PageStruct{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			HasMore:  util.HasMore(int64(len(ruleList)), total, offset),
		},
	}, nil
}

func GetRuleByCondition(ctx context.Context, req *pb_mani.GetRuleByConditionReq) (*pb_mani.GetRuleByConditionResp, error) {
	rules, err := dal.GetRuleByCondition(ctx, req.RuleId, req.User, req.RuleType, req.RuleState, req.ExecuteType)
	if err != nil {
		logx.Errorf("error logic GetRuleByCondition")
		return nil, err
	}
	ruleMap := make(map[string]*pb_mani.Rule, 0)
	for _, r := range rules {
		ruleMap[r.RuleId] = &pb_mani.Rule{
			RuleId:      r.RuleId,
			User:        r.User,
			RuleType:    pb_mani.RuleType(r.RuleType),
			RuleState:   pb_mani.RuleState(r.RuleState),
			ExecuteType: pb_mani.ExecuteType(r.ExecuteType),
			RuleConfig: &pb_mani.RuleConfig{
				DescText:    r.DescText,
				SourceUrl:   r.SourceUrl,
				SourceName:  r.SourceName,
				SourceState: pb_mani.FileState(r.SourceState),
			},
		}
	}
	return &pb_mani.GetRuleByConditionResp{
		Rules: ruleMap,
	}, nil
}

func UpdateRule(ctx context.Context, req *pb_mani.UpdateRuleReq) error {
	updateFiled := make(map[string]interface{}, 0)
	if req.RuleType != pb_mani.RuleType_unknown_rule_type {
		updateFiled["rule_type"] = req.RuleType
	}
	if req.RuleState != pb_mani.RuleState_unknown_rule_state {
		updateFiled["rule_state"] = req.RuleState
	}
	if req.ExecuteType != pb_mani.ExecuteType_unknown_exec_type {
		updateFiled["execute_type"] = req.ExecuteType
	}
	if req.DescText != "" {
		updateFiled["desc_text"] = req.DescText
	}
	if req.SourceUrl != "" {
		updateFiled["source_url"] = req.SourceUrl
	}
	if req.SourceName != "" {
		updateFiled["source_name"] = req.SourceName
	}
	if req.SourceState != pb_mani.FileState_unknown_file_state {
		updateFiled["source_state"] = req.SourceState
	}
	if err := dal.UpdateRule(ctx, req.RuleId, updateFiled); err != nil {
		logx.Errorf("logic UpdateRule error")
		return err
	}
	return nil

}
