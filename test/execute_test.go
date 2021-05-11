package test

import (
	"context"
	"fmt"
	"gin_mani_center/handler"
	pb_mani "gin_mani_center/pb"
	"testing"
)

func TestExecuteRule(t *testing.T) {
	req := &pb_mani.ExecuteRuleReq{
		Rule: &pb_mani.Rule{
			RuleId: "c518ef3f360856b84d971b54623de989",
		},
	}
	resp, err := handler.ExecuteRule(context.Background(), req)
	fmt.Println(resp, err)
}
