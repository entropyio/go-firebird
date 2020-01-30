package db

import (
	"Firebird/config"
	"fmt"
	"time"
)

func (user *UserInfo) TableName() string {
	return "user_info"
}

func QueryUserInfo(userInfoQuery *UserInfoQuery) (count int64, userList []UserInfo) {
	db := DB()
	var resultList []UserInfo
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if userInfoQuery.Id > 0 {
		query.Where("id", "=", userInfoQuery.Id)
	}
	if userInfoQuery.Username != "" {
		query.Where("user_name", "=", userInfoQuery.Username)
	}
	if userInfoQuery.Status > 0 {
		query.Where("status", "=", userInfoQuery.Status)
	}
	if !userInfoQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", userInfoQuery.StartTime)
	}
	if !userInfoQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", userInfoQuery.EndTime)
	}

	if userInfoQuery.PageSize > 0 {
		query.Limit(userInfoQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if userInfoQuery.PageNumber > 0 {
		query.Offset((userInfoQuery.PageNumber - 1) * userInfoQuery.PageSize)
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

func InsertUserInfo(userInfo *UserInfo) (id int64) {
	db := DB()

	userInfo.GmtCreate = time.Now()
	userInfo.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   userInfo.GmtCreate,
		"gmt_modified": userInfo.GmtModified,
		"user_name":    userInfo.Username,
		"user_desc":    userInfo.UserDesc,
		"status":       userInfo.Status,
	}
	query := db.Table(userInfo)
	query.Data(data)

	id, err = query.InsertGetId()
	if nil != err {
		fmt.Println(err)
	}
	return id
}

func UpdateUserInfo(userInfo *UserInfo) (count int64) {
	if userInfo.Id <= 0 {
		return 0
	}

	db := DB()

	userInfo.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": userInfo.GmtModified,
	}
	if userInfo.Username != "" {
		data["user_name"] = userInfo.Username
	}
	if userInfo.UserDesc != "" {
		data["user_desc"] = userInfo.UserDesc
	}
	if userInfo.Status > 0 {
		data["status"] = userInfo.Status
	}

	query := db.Table(userInfo)
	query.Data(data)
	query.Where("id", userInfo.Id)
	count, _ = query.Update()

	return count
}

func DeleteUserInfo(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("user_info")
	query.Where("id", id)
	count, _ = query.Delete()

	return count
}

func GetUserInfoById(id int64) (record UserInfo) {
	if id <= 0 {
		return UserInfo{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("id", id)
	query.Select()

	return record
}

func GetUserInfoByName(name string) (record UserInfo) {
	if "" == name {
		return UserInfo{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("user_name", name)
	query.Where("status", config.STATUS_ENABLE)
	query.Select()

	return record
}

func loadAllUserInfo() (userList []UserInfo) {
	db := DB()
	query := db.Table(&userList)
	query.Where("status", "=", config.STATUS_ENABLE)
	query.Select()
	return userList
}
