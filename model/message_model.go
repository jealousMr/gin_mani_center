package model


type Message struct {
	Id string `json:"id"`
	TaskId string `json:"task_id"`
	RuleId string `json:"rule_id"`
	Operator string `json:"operator"`
	Extra string `json:"extra"`
}