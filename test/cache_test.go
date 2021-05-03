package test

import (
	"fmt"
	"gin_mani_center/dal"
	"gin_mani_center/model"
	"gin_mani_center/util"
	"testing"
)

func TestSendMsg(t *testing.T) {
	m := &model.Message{
		Id:     util.GenUID(),
		TaskId: "test",
		Extra:  "ee",
	}
	err := dal.SendMessage(m)
	fmt.Println(err)
}

func TestGetMeg(t *testing.T) {
	m, err := dal.GetMessage()
	fmt.Println(m, err)
}
