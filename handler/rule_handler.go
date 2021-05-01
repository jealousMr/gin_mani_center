package handler

import (
	"context"
	"errors"
	"gin_mani_center/logic"
	pb_mani "gin_mani_center/pb"
	"gin_mani_center/util"
	logx "github.com/amoghe/distillog"
)

func checkAddRuleParam(req *pb_mani.AddRuleReq) error {
	if req.Rule == nil || req.Rule.RuleConfig == nil{
		return errors.New(util.MsgParamError)
	}
	return nil
}
func AddRule(ctx context.Context, req *pb_mani.AddRuleReq) (resp *pb_mani.AddRuleResp, err error) {
	resp = &pb_mani.AddRuleResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkAddRuleParam(req); err != nil {
		logx.Errorf("AddRule checkAddRuleParam error:%v", err)
		return resp, err
	}
	if err = logic.AddRule(ctx, req); err != nil {
		logx.Errorf("AddRule error:%v", err)
		return resp, err
	}
	return
}

func checkGetRuleByRuleType(req *pb_mani.GetRuleByRuleTypeReq) error {
	if req.RuleType == pb_mani.RuleType_unknown_rule_type {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func GetRuleByRuleType(ctx context.Context, req *pb_mani.GetRuleByRuleTypeReq) (resp *pb_mani.GetRuleByRuleTypeResp, err error) {
	resp = &pb_mani.GetRuleByRuleTypeResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err := checkGetRuleByRuleType(req); err != nil {
		logx.Errorf("GetRuleByRuleType checkGetRuleByRuleType error:%v", err)
		return resp, err
	}
	resp, err = logic.GetRuleByRuleType(ctx, req)
	if err != nil {
		logx.Errorf("GetRuleByRuleType  error:%v", err)
		return resp, err
	}
	return resp, err
}

func GetRuleByCondition(ctx context.Context, req *pb_mani.GetRuleByConditionReq) (resp *pb_mani.GetRuleByConditionResp, err error) {
	resp = &pb_mani.GetRuleByConditionResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	resp, err = logic.GetRuleByCondition(ctx, req)
	if err != nil {
		logx.Errorf("GetRuleByCondition  error:%v", err)
		return resp, err
	}
	return
}

func checkUpdateRule(req *pb_mani.UpdateRuleReq) error {
	if req.RuleId == "" {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func UpdateRule(ctx context.Context, req *pb_mani.UpdateRuleReq) (resp *pb_mani.UpdateRuleResp, err error) {
	resp = &pb_mani.UpdateRuleResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err := checkUpdateRule(req); err != nil {
		logx.Errorf("UpdateRule checkUpdateRule  error:%v", err)
		return resp, err
	}
	if err := logic.UpdateRule(ctx, req); err != nil {
		logx.Errorf("UpdateRule  error:%v", err)
		return resp, err
	}
	return

}
