package clients

import (
	"fmt"
	"gin_mani_center/conf"
	pb_mani "gin_mani_center/pb"
	"google.golang.org/grpc"
	"log"
)

var engineClient pb_mani.GinEngineServiceClient

func GetEngineClient() (pb_mani.GinEngineServiceClient, error) {
	if engineClient != nil {
		return engineClient, nil
	}
	var err error
	engineClient, err = connectEngine()
	return engineClient, err
}

func connectEngine() (pb_mani.GinEngineServiceClient, error) {
	cf := conf.GetConfig()
	address := fmt.Sprintf("localhost%s", cf.Client.Engine)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("engine service client connect err: ", err)
	}
	c := pb_mani.NewGinEngineServiceClient(conn)
	return c, err
}