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
			RuleId: "19e3ca476e516df01f3c814b042ee105",
		},
	}
	resp, err := handler.ExecuteRule(context.Background(), req)
	fmt.Println(resp, err)
}
