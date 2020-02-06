package db

import (
	"Firebird/config"
	"time"
)

func (user *UserSchedule) TableName() string {
	return "user_schedule"
}

func QueryUserSchedule(userScheduleQuery *UserScheduleQuery) (count int64, userList []UserSchedule) {
	db := DB()
	var resultList []UserSchedule
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if userScheduleQuery.Id > 0 {
		query.Where("id", "=", userScheduleQuery.Id)
	}
	if userScheduleQuery.UserId > 0 {
		query.Where("user_id", "=", userScheduleQuery.UserId)
	}
	if userScheduleQuery.SymbolId > 0 {
		query.Where("symbol_id", "=", userScheduleQuery.SymbolId)
	}
	if userScheduleQuery.Type > 0 {
		query.Where("type", "=", userScheduleQuery.Type)
	}
	if userScheduleQuery.Status > 0 {
		query.Where("status", "=", userScheduleQuery.Status)
	}
	if !userScheduleQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", userScheduleQuery.StartTime)
	}
	if !userScheduleQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", userScheduleQuery.EndTime)
	}

	if userScheduleQuery.PageSize > 0 {
		query.Limit(userScheduleQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if userScheduleQuery.PageNumber > 0 {
		query.Offset((userScheduleQuery.PageNumber - 1) * userScheduleQuery.PageSize)
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

func InsertUserSchedule(userSchedule *UserSchedule) (id int64) {
	db := DB()

	userSchedule.GmtCreate = time.Now()
	userSchedule.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   userSchedule.GmtCreate,
		"gmt_modified": userSchedule.GmtModified,
		"user_id":      userSchedule.UserId,
		"symbol_id":    userSchedule.SymbolId,
		"name":         userSchedule.Name,
		"type":         userSchedule.Type,
		"amount":       userSchedule.Amount,
		"success":      userSchedule.Success,
		"failed":       userSchedule.Failed,
		"status":       userSchedule.Status,
	}
	query := db.Table(userSchedule)
	query.Data(data)
	id, err = query.InsertGetId()
	if nil != err {
		log.Error(err)
	}
	userSchedule.Id = id
	return id
}

func UpdateUserSchedule(userSchedule *UserSchedule) (count int64) {
	if userSchedule.Id <= 0 {
		return 0
	}

	db := DB()

	userSchedule.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": userSchedule.GmtModified,
	}
	if userSchedule.Name != "" {
		data["name"] = userSchedule.Name
	}
	if userSchedule.UserId > 0 {
		data["user_id"] = userSchedule.UserId
	}
	if userSchedule.SymbolId > 0 {
		data["symbol_id"] = userSchedule.SymbolId
	}
	if userSchedule.Type > 0 {
		data["type"] = userSchedule.Type
	}
	if userSchedule.Amount > 0 {
		data["amount"] = userSchedule.Amount
	}
	if userSchedule.Success > 0 {
		data["success"] = userSchedule.Success
	}
	if userSchedule.Failed > 0 {
		data["failed"] = userSchedule.Failed
	}
	if userSchedule.Status > 0 {
		data["status"] = userSchedule.Status
	}

	query := db.Table(userSchedule)
	query.Data(data)
	query.Where("id", userSchedule.Id)
	count, err = query.Update()
	if nil != err {
		log.Error(err)
	}

	return count
}

func DeleteUserSchedule(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("user_schedule")
	query.Where("id", id)
	count, err = query.Delete()
	if nil != err {
		log.Error(err)
	}
	return count
}

func UpdateScheduleRule(userSchedule *UserSchedule, ruleList []RuleItem) {
	deleteScheduleRules(userSchedule.Id)
	for i := range ruleList {
		ruleItem := ruleList[i]
		ruleItem.Id = 0
		ruleItem.UserId = userSchedule.UserId
		ruleItem.SymbolId = userSchedule.SymbolId
		ruleItem.ScheduleId = userSchedule.Id
		ruleItem.Status = config.STATUS_ENABLE
		InsertRuleItem(&ruleItem)
	}
}

func loadAllSchedule() (resultList []UserSchedule) {
	db := DB()
	query := db.Table(&resultList)
	query.Where("status", "=", config.STATUS_ENABLE)
	err := query.Select()
	if nil != err {
		log.Error(err)
	}
	return resultList
}
