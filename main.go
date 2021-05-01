package main

import (
	"context"
	"gin_mani_center/conf"
	"gin_mani_center/handler"
	pb_mani "gin_mani_center/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	logx "github.com/amoghe/distillog"
	"log"
	"net"
)

type S struct {
}

func (s S) AddRule(ctx context.Context, req *pb_mani.AddRuleReq) (*pb_mani.AddRuleResp, error) {
	logx.Infoln("AddRule req: %v", req)
	resp, err := handler.AddRule(ctx, req)
	logx.Infoln("AddRule resp: %v", resp)
	return resp, err
}

func (s S) GetRuleByRuleType(ctx context.Context, req *pb_mani.GetRuleByRuleTypeReq) (*pb_mani.GetRuleByRuleTypeResp, error) {
	logx.Infoln("GetRuleByRuleType req: %v", req)
	resp, err := handler.GetRuleByRuleType(ctx, req)
	logx.Infoln("GetRuleByRuleType resp: %v", resp)
	return resp, err
}

func (s S) GetRuleByCondition(ctx context.Context, req *pb_mani.GetRuleByConditionReq) (*pb_mani.GetRuleByConditionResp, error) {
	logx.Infoln("GetRuleByCondition req: %v", req)
	resp, err := handler.GetRuleByCondition(ctx, req)
	logx.Infoln("GetRuleByCondition resp: %v", resp)
	return resp, err
}

func (s S) UpdateRule(ctx context.Context, req *pb_mani.UpdateRuleReq) (*pb_mani.UpdateRuleResp, error) {
	logx.Infoln("UpdateRule req: %v", req)
	resp, err := handler.UpdateRule(ctx, req)
	logx.Infoln("UpdateRule resp: %v", resp)
	return resp, err
}

func main() {
	cf := conf.GetConfig()
	logx.Infof("start mani center server")
	lis, err := net.Listen("tcp", cf.Server.Ip)
	if err != nil {
		log.Fatal("failed to listen")
	}
	server := grpc.NewServer()
	pb_mani.RegisterGinCenterServiceServer(server, &S{})
	reflection.Register(server)
	logx.Infof("run mani center server success...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
