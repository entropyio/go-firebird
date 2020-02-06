package db

import (
	"Firebird/config"
	"time"
)

func (rule *RuleItem) TableName() string {
	return "rule_item"
}

func QueryRuleItem(ruleInfoQuery *RuleItemQuery) (count int64, ruleList []RuleItem) {
	db := DB()
	var resultList []RuleItem
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if ruleInfoQuery.Id > 0 {
		query.Where("id", "=", ruleInfoQuery.Id)
	}
	if ruleInfoQuery.UserId > 0 {
		query.Where("user_id", "=", ruleInfoQuery.UserId)
	}
	if ruleInfoQuery.SymbolId > 0 {
		query.Where("symbol_id", "=", ruleInfoQuery.SymbolId)
	}
	if ruleInfoQuery.ScheduleId > 0 {
		query.Where("schedule_id", "=", ruleInfoQuery.ScheduleId)
	}

	if ruleInfoQuery.PageSize > 0 {
		query.Limit(ruleInfoQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if ruleInfoQuery.PageNumber > 0 {
		query.Offset((ruleInfoQuery.PageNumber - 1) * ruleInfoQuery.PageSize)
	} else {
		query.Offset(0)
	}
	// get count
	resultCount, err = query.Count("id")
	err = query.Select()
	if nil != err {
		log.Error(err)
	}

	return resultCount, resultList
}

func InsertRuleItem(ruleItem *RuleItem) (id int64) {
	db := DB()

	ruleItem.GmtCreate = time.Now()
	ruleItem.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   ruleItem.GmtCreate,
		"gmt_modified": ruleItem.GmtModified,
		"user_id":      ruleItem.UserId,
		"symbol_id":    ruleItem.SymbolId,
		"schedule_id":  ruleItem.ScheduleId,
		"rule_type":    ruleItem.RuleType,
		"join_type":    ruleItem.JoinType,
		"op_type":      ruleItem.OpType,
		"op_value":     ruleItem.OpValue,
		"status":       ruleItem.Status,
	}
	query := db.Table(ruleItem)
	query.Data(data)
	id, err = query.InsertGetId()
	if nil != err {
		log.Error(err)
	}
	ruleItem.Id = id
	return id
}

func UpdateRuleItem(ruleItem *RuleItem) (count int64) {
	if ruleItem.Id <= 0 {
		return 0
	}

	db := DB()

	ruleItem.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": ruleItem.GmtModified,
	}
	if ruleItem.UserId > 0 {
		data["user_id"] = ruleItem.UserId
	}
	if ruleItem.SymbolId > 0 {
		data["symbol_id"] = ruleItem.SymbolId
	}
	if ruleItem.ScheduleId > 0 {
		data["schedule_id"] = ruleItem.ScheduleId
	}
	if ruleItem.RuleType > 0 {
		data["rule_type"] = ruleItem.RuleType
	}
	if ruleItem.JoinType > 0 {
		data["join_type"] = ruleItem.JoinType
	}
	if ruleItem.OpType > 0 {
		data["op_type"] = ruleItem.OpType
	}
	if ruleItem.OpValue != "" {
		data["op_value"] = ruleItem.OpValue
	}
	if ruleItem.Status > 0 {
		data["status"] = ruleItem.Status
	}

	query := db.Table(ruleItem)
	query.Data(data)
	query.Where("id", ruleItem.Id)
	count, err = query.Update()
	if nil != err {
		log.Error(err)
	}

	return count
}

func DeleteRuleItem(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("rule_item")
	query.Where("id", id)
	count, err = query.Delete()
	if nil != err {
		log.Error(err)
	}

	return count
}

func deleteScheduleRules(scheduleId int64) (count int64) {
	if scheduleId <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("rule_item")
	query.Where("schedule_id", scheduleId)
	count, err = query.Delete()
	if nil != err {
		log.Error(err)
	}

	return count
}

func loadAllRuleItem() (ruleList []RuleItem) {
	db := DB()
	query := db.Table(&ruleList)
	query.Where("status", "=", config.STATUS_ENABLE)
	err := query.Select()
	if nil != err {
		log.Error(err)
	}
	return ruleList
}
