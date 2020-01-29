package db

import (
	"time"
)

func (user *UserAccount) TableName() string {
	return "user_account"
}

func QueryUserAccount(userAccountQuery *UserAccountQuery) (count int64, userList []UserAccount) {
	db := DB()
	var resultList []UserAccount
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if userAccountQuery.Id > 0 {
		query.Where("id", "=", userAccountQuery.Id)
	}
	if userAccountQuery.UserId > 0 {
		query.Where("user_id", "=", userAccountQuery.UserId)
	}
	if userAccountQuery.SymbolId > 0 {
		query.Where("symbol_id", "=", userAccountQuery.SymbolId)
	}
	if userAccountQuery.Status > 0 {
		query.Where("status", "=", userAccountQuery.Status)
	}
	if !userAccountQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", userAccountQuery.StartTime)
	}
	if !userAccountQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", userAccountQuery.EndTime)
	}

	if userAccountQuery.PageSize > 0 {
		query.Limit(userAccountQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if userAccountQuery.PageNumber > 0 {
		query.Offset((userAccountQuery.PageNumber - 1) * userAccountQuery.PageSize)
	} else {
		query.Offset(0)
	}
	// get count
	resultCount, err = query.Count("id")
	// output sql
	//sql, _, _ := query.BuildSql()
	//fmt.Println(sql)
	// run query
	query.Select()

	return resultCount, resultList
}

func InsertUserAccount(userAccount *UserAccount) (id int64) {
	db := DB()

	userAccount.GmtCreate = time.Now()
	userAccount.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":    userAccount.GmtCreate,
		"gmt_modified":  userAccount.GmtModified,
		"user_id":       userAccount.UserId,
		"symbol_id":     userAccount.SymbolId,
		"hold_amount":   userAccount.HoldAmount,
		"hold_price":    userAccount.HoldPrice,
		"yest_benifit":  userAccount.YestBenefit,
		"total_benifit": userAccount.TotalBenefit,
		"status":        userAccount.Status,
	}
	query := db.Table(userAccount)
	query.Data(data)
	// output sql
	//sql, _, _ := query.BuildSql()
	//fmt.Println(sql)

	id, _ = query.InsertGetId()
	userAccount.Id = id

	return id
}

func UpdateUserAccount(userAccount *UserAccount) (count int64) {
	if userAccount.Id <= 0 {
		return 0
	}

	db := DB()

	userAccount.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": userAccount.GmtModified,
	}
	if userAccount.UserId > 0 {
		data["user_id"] = userAccount.UserId
	}
	if userAccount.SymbolId > 0 {
		data["symbol_id"] = userAccount.SymbolId
	}

	data["price"] = userAccount.Price
	data["amount"] = userAccount.Amount
	data["total"] = userAccount.Total
	data["rate"] = userAccount.Rate
	data["hold_amount"] = userAccount.HoldAmount
	data["hold_price"] = userAccount.HoldPrice
	data["yest_benefit"] = userAccount.YestBenefit
	data["total_benefit"] = userAccount.TotalBenefit

	if userAccount.Status > 0 {
		data["status"] = userAccount.Status
	}

	query := db.Table(userAccount)
	query.Data(data)
	query.Where("id", userAccount.Id)
	count, _ = query.Update()

	return count
}

func DeleteUserAccount(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("user_account")
	query.Where("id", id)
	count, _ = query.Delete()

	return count
}

func GetUserAccountById(id int64) (record UserAccount) {
	if id <= 0 {
		return UserAccount{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("id", id)
	query.Select()

	return record
}

func GetUserAccountByUid(userId int64, symbolId int64) (record UserAccount) {
	if userId <= 0 || symbolId <= 0 {
		return UserAccount{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("user_id", userId)
	query.Where("symbol_id", symbolId)
	query.Select()

	return record
}
