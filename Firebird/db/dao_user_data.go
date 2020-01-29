package db

import (
	"time"
)

func (trade *UserData) TableName() string {
	return "user_data"
}

func QueryUserData(userDataQuery *UserDataQuery) (count int64, dataList []UserData) {
	db := DB()
	var resultList []UserData
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("gmt_create DESC")
	if userDataQuery.Id > 0 {
		query.Where("id", "=", userDataQuery.Id)
	}
	if userDataQuery.UserId > 0 {
		query.Where("user_id", "=", userDataQuery.UserId)
	}
	if userDataQuery.SymbolId >= 0 {
		query.Where("symbol_id", "=", userDataQuery.SymbolId)
	}
	if userDataQuery.Status > 0 {
		query.Where("status", "=", userDataQuery.Status)
	}
	if !userDataQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", userDataQuery.StartTime)
	}
	if !userDataQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", userDataQuery.EndTime)
	}

	if userDataQuery.PageSize > 0 {
		query.Limit(userDataQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if userDataQuery.PageNumber > 0 {
		query.Offset((userDataQuery.PageNumber - 1) * userDataQuery.PageSize)
	} else {
		query.Offset(0)
	}
	// get count
	resultCount, err = query.Count("id")
	query.Select()

	return resultCount, resultList
}

func InsertUserData(userData *UserData) (id int64) {
	db := DB()
	if userData.GmtCreate.IsZero() {
		userData.GmtCreate = time.Now()
		userData.GmtModified = time.Now()
	}

	var data = map[string]interface{}{
		"gmt_create":   userData.GmtCreate,
		"gmt_modified": userData.GmtModified,
		"user_id":      userData.UserId,
		"symbol_id":    userData.SymbolId,
		"open_price":   userData.OpenPrice,
		"close_price":  userData.ClosePrice,
		"high_price":   userData.HighPrice,
		"low_price":    userData.LowPrice,
		"hold_price":   userData.HoldPrice,
		"hold_amount":  userData.HoldAmount,
		"hold_rate":    userData.HoldRate,
		"hold_benefit": userData.HoldBenefit,
		"status":       userData.Status,
	}
	query := db.Table(userData)
	query.Data(data)

	id, _ = query.InsertGetId()
	userData.Id = id
	return id
}

func UpdateUserData(userData *UserData) (count int64) {
	if userData.Id <= 0 {
		return 0
	}

	db := DB()
	userData.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_modified": userData.GmtModified,
	}
	if userData.UserId > 0 {
		data["user_id"] = userData.UserId
	}
	if userData.SymbolId > 0 {
		data["symbol_id"] = userData.SymbolId
	}
	if userData.OpenPrice != 0 {
		data["open_price"] = userData.OpenPrice
	}
	if userData.ClosePrice != 0 {
		data["close_price"] = userData.ClosePrice
	}
	if userData.HighPrice != 0 {
		data["high_price"] = userData.HighPrice
	}
	if userData.LowPrice != 0 {
		data["low_price"] = userData.LowPrice
	}
	if userData.HoldPrice != 0 {
		data["hold_price"] = userData.HoldPrice
	}
	if userData.HoldAmount != 0 {
		data["hold_amount"] = userData.HoldAmount
	}
	if userData.HoldRate != 0 {
		data["hold_rate"] = userData.HoldRate
	}
	if userData.HoldBenefit != 0 {
		data["hold_benefit"] = userData.HoldBenefit
	}
	if userData.Status > 0 {
		data["status"] = userData.Status
	}

	query := db.Table(userData)
	query.Data(data)
	query.Where("id", userData.Id)
	count, _ = query.Update()

	return count
}

func DeleteUserData(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("user_data")
	query.Where("id", id)
	count, _ = query.Delete()

	return count
}
