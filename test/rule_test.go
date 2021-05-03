package test

import (
	"context"
	"fmt"
	"gin_mani_center/conf"
	"gin_mani_center/dal"
	"gin_mani_center/handler"
	pb_mani "gin_mani_center/pb"
	"testing"
)

func TestAddRule(t *testing.T) {
	req := &pb_mani.AddRuleReq{
		Rule: &pb_mani.Rule{
			User: "3219jdsia8923",
			RuleConfig: &pb_mani.RuleConfig{
				SourceName: "ee",
				SourceState: pb_mani.FileState_file_invalid,
			},
			RuleType: pb_mani.RuleType_default_all,
			RuleState: pb_mani.RuleState_rule_state_valid,
		},
	}
	resp, err := handler.AddRule(context.Background(), req)
	fmt.Println(resp, err)
}

func TestConfig(t *testing.T) {
	fmt.Println(conf.GetConfig())
}

func TestGetRuleByIdDal(t *testing.T) {
	rule, err := dal.GetRuleById(context.Background(), "c518ef3f360856b84d971b54623de989")
	fmt.Println(rule, err)
}
