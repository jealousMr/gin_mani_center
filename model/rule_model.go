package model

type RuleModel struct {
	RuleId string `json:"rule_id"`
	User string `json:"user"`
	RuleType int64 `json:"rule_type"`
	RuleState int64 `json:"rule_state"`
	ExecuteType int64 `json:"execute_type"`
	DescText string `json:"desc_text"`
	SourceName string `json:"source_name"`
	SourceUrl string `json:"source_url"`
	SourceState int64 `json:"source_state"`
}

func RuleTableName() string{
	return "rule"
}

func (RuleModel) TableName() string{
	return RuleTableName()
}
