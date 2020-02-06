package db

import (
	"time"
)

func (trade *UserTrade) TableName() string {
	return "user_trade"
}

func QueryUserTrade(userTradeQuery *UserTradeQuery) (count int64, tradeList []UserTrade) {
	db := DB()
	var resultList []UserTrade
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("gmt_create DESC")
	if userTradeQuery.Id > 0 {
		query.Where("id", "=", userTradeQuery.Id)
	}
	if userTradeQuery.UserId > 0 {
		query.Where("user_id", "=", userTradeQuery.UserId)
	}
	if userTradeQuery.SymbolId > 0 {
		query.Where("symbol_id", "=", userTradeQuery.SymbolId)
	}
	if userTradeQuery.Type > 0 {
		query.Where("type", "=", userTradeQuery.Type)
	}
	if userTradeQuery.ScheduleId > 0 {
		query.Where("schedule_id", "=", userTradeQuery.ScheduleId)
	}
	if userTradeQuery.Status > 0 {
		query.Where("status", "=", userTradeQuery.Status)
	}
	if !userTradeQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", userTradeQuery.StartTime)
	}
	if !userTradeQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", userTradeQuery.EndTime)
	}

	if userTradeQuery.PageSize > 0 {
		query.Limit(userTradeQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if userTradeQuery.PageNumber > 0 {
		query.Offset((userTradeQuery.PageNumber - 1) * userTradeQuery.PageSize)
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

func InsertUserTrade(userTrade *UserTrade) (id int64) {
	db := DB()
	userTrade.GmtCreate = time.Now()
	userTrade.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   userTrade.GmtCreate,
		"gmt_modified": userTrade.GmtModified,
		"user_id":      userTrade.UserId,
		"symbol_id":    userTrade.SymbolId,
		"price":        userTrade.Price,
		"amount":       userTrade.Amount,
		"type":         userTrade.Type,
		"schedule_id":  userTrade.ScheduleId,
		"hold_price":   userTrade.HoldPrice,
		"hold_amount":  userTrade.HoldAmount,
		"status":       userTrade.Status,
	}
	query := db.Table(userTrade)
	query.Data(data)

	id, err = query.InsertGetId()
	if nil != err {
		log.Error(err)
	}
	userTrade.Id = id
	return id
}

func UpdateUserTrade(userTrade *UserTrade) (count int64) {
	if userTrade.Id <= 0 {
		return 0
	}

	db := DB()
	userTrade.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_modified": userTrade.GmtModified,
	}
	if userTrade.UserId > 0 {
		data["user_id"] = userTrade.UserId
	}
	if userTrade.SymbolId > 0 {
		data["symbol_id"] = userTrade.SymbolId
	}
	if userTrade.Price != 0 {
		data["price"] = userTrade.Price
	}
	if userTrade.Amount != 0 {
		data["amount"] = userTrade.Amount
	}
	if userTrade.Type != 0 {
		data["type"] = userTrade.Type
	}
	if userTrade.ScheduleId != 0 {
		data["schedule_id"] = userTrade.ScheduleId
	}
	if userTrade.HoldPrice != 0 {
		data["hold_price"] = userTrade.HoldPrice
	}
	if userTrade.HoldAmount != 0 {
		data["hold_amount"] = userTrade.HoldAmount
	}
	if userTrade.Status > 0 {
		data["status"] = userTrade.Status
	}

	query := db.Table(userTrade)
	query.Data(data)
	query.Where("id", userTrade.Id)
	count, err = query.Update()
	if nil != err {
		log.Error(err)
	}
	return count
}

func DeleteUserTrade(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("user_trade")
	query.Where("id", id)
	count, err = query.Delete()
	if nil != err {
		log.Error(err)
	}
	return count
}
