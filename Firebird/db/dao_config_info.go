package db

import (
	"Firebird/config"
	"time"
)

func (config *ConfigInfo) TableName() string {
	return "config_info"
}

func QueryConfigInfo(configInfoQuery *ConfigInfoQuery) (count int64, configList []ConfigInfo) {
	db := DB()
	var resultList []ConfigInfo
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if configInfoQuery.Id > 0 {
		query.Where("id", "=", configInfoQuery.Id)
	}
	if configInfoQuery.Ckey != "" {
		query.Where("ckey", "like", "%"+configInfoQuery.Ckey+"%")
	}
	if configInfoQuery.Status > 0 {
		query.Where("status", "=", configInfoQuery.Status)
	}

	if configInfoQuery.PageSize > 0 {
		query.Limit(configInfoQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if configInfoQuery.PageNumber > 0 {
		query.Offset((configInfoQuery.PageNumber - 1) * configInfoQuery.PageSize)
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

func InsertConfigInfo(configInfo *ConfigInfo) (id int64) {
	db := DB()

	configInfo.GmtCreate = time.Now()
	configInfo.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   configInfo.GmtCreate,
		"gmt_modified": configInfo.GmtModified,
		"ckey":    configInfo.Ckey,
		"cvalue":    configInfo.Cvalue,
		"operator":    configInfo.Operator,
		"creator":    configInfo.Creator,
		"status":       configInfo.Status,
	}
	query := db.Table(configInfo)
	query.Data(data)

	id, err = query.InsertGetId()
	if nil != err {
		log.Error(err)
	}
	configInfo.Id = id
	return id
}

func UpdateConfigInfo(configInfo *ConfigInfo) (count int64) {
	if configInfo.Id <= 0 {
		return 0
	}

	db := DB()

	configInfo.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": configInfo.GmtModified,
	}
	if configInfo.Ckey != "" {
		data["ckey"] = configInfo.Ckey
	}
	if configInfo.Cvalue != "" {
		data["cvalue"] = configInfo.Cvalue
	}
	if configInfo.Operator != "" {
		data["operator"] = configInfo.Operator
	}
	if configInfo.Creator != "" {
		data["creator"] = configInfo.Creator
	}
	if configInfo.Status > 0 {
		data["status"] = configInfo.Status
	}

	query := db.Table(configInfo)
	query.Data(data)
	query.Where("id", configInfo.Id)
	count, err = query.Update()
	if nil != err {
		log.Error(err)
	}
	return count
}

func DeleteConfigInfo(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("config_info")
	query.Where("id", id)
	count, err = query.Delete()
	if nil != err {
		log.Error(err)
	}
	return count
}

func GetConfigInfoById(id int64) (record ConfigInfo) {
	if id <= 0 {
		return ConfigInfo{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("id", id)
	err := query.Select()
	if nil != err {
		log.Error(err)
	}
	return record
}

func GetConfigInfoByKey(key string) (record ConfigInfo) {
	if "" == key {
		return ConfigInfo{}
	}

	db := DB()
	query := db.Table(&record)
	query.Where("ckey", key)
	query.Where("status", config.STATUS_ENABLE)
	err := query.Select()
	if nil != err {
		log.Error(err)
	}
	return record
}

func loadAllConfigInfo() (configList []ConfigInfo) {
	db := DB()
	query := db.Table(&configList)
	query.Where("status", "=", config.STATUS_ENABLE)
	err := query.Select()
	if nil != err {
		log.Error(err)
	}
	return configList
}
