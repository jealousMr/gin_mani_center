package dal

import (
	"bytes"
	"encoding/json"
	"gin_mani_center/model"
	logx "github.com/amoghe/distillog"
)

const QueueKey = "task_msg_queue_flag_key"

func SendMessage(msg *model.Message) error {
	rd := GetRedisProxy().Get()
	by, err := json.Marshal(msg)
	if err != nil {
		logx.Errorf("SendMessage Marshal error:%v", err)
		return err
	}
	_, err = rd.Do("rpush", QueueKey, by)
	if err != nil {
		logx.Errorf("SendMessage error:%v", err)
		return err
	}
	return nil
}

func GetMessage() (*model.Message, error) {
	rd := GetRedisProxy().Get()
	msg, err := rd.Do("LPOP", QueueKey)
	if err != nil {
		logx.Errorf("GetMessage error:%v", err)
		return nil, err
	}
	if msg == nil {
		logx.Warningln("GetMessage current no message")
		return nil, nil
	}
	message := &model.Message{}
	decoder := json.NewDecoder(bytes.NewReader(msg.([]byte)))
	if err := decoder.Decode(&message); err != nil {
		logx.Errorf("GetMessage Decode error:%v", err)
		return nil, err
	}
	return message, nil
}