package test

import (
	"context"
	"fmt"
	"gin_mani_center/handler"
	pb_mani "gin_mani_center/pb"
	"testing"
)

func TestAddRule(t *testing.T) {
	req := &pb_mani.AddRuleReq{
		Rule: &pb_mani.Rule{
			User: "sysall",
			RuleConfig: &pb_mani.RuleConfig{
				SourceName: "ee",
			},
			RuleType: pb_mani.RuleType_default_all,
		},
	}
	resp, err := handler.AddRule(context.Background(), req)
	fmt.Println(resp, err)
}
