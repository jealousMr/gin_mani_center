package dal

import (
	"context"
	"gin_mani_center/model"
	pb_mani "gin_mani_center/pb"
	logx "github.com/amoghe/distillog"
)

func AddRule(ctx context.Context, rule *model.RuleModel) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return err
	}
	if err := db.Create(&rule).Error; err != nil {
		logx.Errorf("dal AddRule error:%v", err)
		return err
	}
	return nil
}

func GetRuleByRuleType(ctx context.Context, ruleType pb_mani.RuleType, limit, offset int) (ruleList []*model.RuleModel, total int64, err error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return nil, 0, err
	}
	ruleList = make([]*model.RuleModel, 0)
	if err = db.Table(model.RuleTableName()).Offset(offset).Limit(limit).Find(&ruleList).Error; err != nil {
		logx.Errorf("dal GetRuleByRuleType error:%v", err)
		return nil, 0, err
	}
	if err = db.Table(model.RuleTableName()).Count(&total).Error; err != nil {
		logx.Errorf("dal GetRuleByRuleType count error:%v", err)
		return nil, 0, err
	}
	return ruleList, total, nil

}

func GetRuleByCondition(ctx context.Context, ruleId, userId string, ruleType pb_mani.RuleType, ruleState pb_mani.RuleState, executeType pb_mani.ExecuteType) ([]*model.RuleModel, error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return nil, err
	}
	ruleList := make([]*model.RuleModel, 0)
	db = db.Table(model.RuleTableName())
	if ruleId != "" {
		db = db.Where("rule_id = ?", ruleId)
	}
	if userId != "" {
		db = db.Where("user_id = ?", userId)
	}
	if ruleType != pb_mani.RuleType_unknown_rule_type {
		db = db.Where("rule_type = ?", ruleType)
	}
	if ruleState != pb_mani.RuleState_unknown_rule_state {
		db = db.Where("rule_state = ?", ruleState)
	}
	if executeType != pb_mani.ExecuteType_unknown_exec_type {
		db = db.Where("execute_type = ?", executeType)
	}
	if err := db.Find(&ruleList).Error; err != nil {
		logx.Errorf("dal GetRuleByCondition error:%v", err)
		return nil, err
	}
	return ruleList, nil
}

func UpdateRule(ctx context.Context, ruleId string, updateFiled map[string]interface{}) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return err
	}
	if err := db.Table(model.RuleTableName()).Where("rule_id = ?", ruleId).Updates(updateFiled).Error; err != nil {
		logx.Errorf("dal UpdateRule error:%v", err)
		return err
	}
	return nil
}
