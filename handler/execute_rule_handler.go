package handler

import (
	"context"
	"errors"
	"gin_mani_center/logic"
	pb_mani "gin_mani_center/pb"
	"gin_mani_center/util"
	logx "github.com/amoghe/distillog"
)

func checkExecuteRule(req *pb_mani.ExecuteRuleReq)error{
	if req.Rule == nil || req.Rule.RuleConfig == nil{
		return errors.New(util.MsgParamError)
	}
	return nil
}

func ExecuteRule(ctx context.Context, req *pb_mani.ExecuteRuleReq) (resp *pb_mani.ExecuteRuleResp, err error) {
	resp = &pb_mani.ExecuteRuleResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkExecuteRule(req);err != nil{
		logx.Errorf("ExecuteRule checkExecuteRuleParam error:%v", err)
		return resp, err
	}
	taskId, err := logic.ExecuteRule(ctx, req.Rule)
	if err != nil {
		logx.Errorf("ExecuteRule  error:%v", err)
		return resp, err
	}
	resp.TaskId = taskId
	return
}
