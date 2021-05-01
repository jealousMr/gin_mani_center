package test

import (
	"context"
	"fmt"
	"gin_mani_center/handler"
	pb_mani "gin_mani_center/pb"
	"testing"
)

func TestGetRuleByRuleType(t *testing.T) {
	req := &pb_mani.GetRuleByRuleTypeReq{
		RuleType: pb_mani.RuleType_default_all,
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := handler.GetRuleByRuleType(context.Background(), req)
	fmt.Println(resp, err)
}

func TestGetRuleByCondition(t *testing.T) {
	req := &pb_mani.GetRuleByConditionReq{
		RuleType: pb_mani.RuleType_default_all,
	}
	resp, err := handler.GetRuleByCondition(context.Background(), req)
	fmt.Println(resp, err)
}

func TestUpdateRule(t *testing.T) {
	req := &pb_mani.UpdateRuleReq{
		RuleId:   "21019476f7626b394c2bf2ee76bf8e05",
		RuleType: pb_mani.RuleType_default_image,
	}
	resp, err := handler.UpdateRule(context.Background(), req)
	fmt.Println(resp, err)
}
