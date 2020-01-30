package db

import (
	"Firebird/config"
	"time"
)

func (symbol *SymbolInfo) TableName() string {
	return "symbol_info"
}

func QuerySymbolInfo(symbolInfoQuery *SymbolInfoQuery) (count int64, userList []SymbolInfo) {
	db := DB()
	var resultList []SymbolInfo
	var resultCount int64

	query := db.Table(&resultList)
	query.OrderBy("id DESC")
	if symbolInfoQuery.Id > 0 {
		query.Where("id", "=", symbolInfoQuery.Id)
	}
	if symbolInfoQuery.SymbolName != "" {
		query.Where("symbol_name", "=", symbolInfoQuery.SymbolName)
	}
	if symbolInfoQuery.Status > 0 {
		query.Where("status", "=", symbolInfoQuery.Status)
	}
	if !symbolInfoQuery.StartTime.IsZero() {
		query.Where("gmt_create", ">=", symbolInfoQuery.StartTime)
	}
	if !symbolInfoQuery.EndTime.IsZero() {
		query.Where("gmt_create", "<=", symbolInfoQuery.EndTime)
	}

	if symbolInfoQuery.PageSize > 0 {
		query.Limit(symbolInfoQuery.PageSize)
	} else {
		query.Limit(10)
	}

	if symbolInfoQuery.PageNumber > 0 {
		query.Offset((symbolInfoQuery.PageNumber - 1) * symbolInfoQuery.PageSize)
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

func InsertSymbolInfo(symbolInfo *SymbolInfo) (id int64) {
	db := DB()

	symbolInfo.GmtCreate = time.Now()
	symbolInfo.GmtModified = time.Now()

	var data = map[string]interface{}{
		"gmt_create":   symbolInfo.GmtCreate,
		"gmt_modified": symbolInfo.GmtModified,
		"symbol_name":  symbolInfo.SymbolName,
		"symbol_desc":  symbolInfo.SymbolDesc,
		"symbol_icon":  symbolInfo.SymbolIcon,
		"symbol_group": symbolInfo.SymbolGroup,
		"status":       symbolInfo.Status,
	}
	query := db.Table(symbolInfo)
	query.Data(data)
	// output sql
	//sql, _, _ := query.BuildSql()
	//fmt.Println(sql)

	id, _ = query.InsertGetId()
	symbolInfo.Id = id

	return id
}

func UpdateSymbolInfo(symbolInfo *SymbolInfo) (count int64) {
	if symbolInfo.Id <= 0 {
		return 0
	}

	db := DB()

	symbolInfo.GmtModified = time.Now()
	var data = map[string]interface{}{
		"gmt_modified": symbolInfo.GmtModified,
	}
	if symbolInfo.SymbolName != "" {
		data["symbol_name"] = symbolInfo.SymbolName
	}
	if symbolInfo.SymbolDesc != "" {
		data["symbol_desc"] = symbolInfo.SymbolDesc
	}
	if symbolInfo.SymbolIcon != "" {
		data["symbol_icon"] = symbolInfo.SymbolIcon
	}
	if symbolInfo.SymbolGroup != "" {
		data["symbol_group"] = symbolInfo.SymbolGroup
	}
	if symbolInfo.Status > 0 {
		data["status"] = symbolInfo.Status
	}

	query := db.Table(symbolInfo)
	query.Data(data)
	query.Where("id", symbolInfo.Id)
	count, _ = query.Update()

	// output sql
	//sql, _, _ := query.BuildSql()
	//fmt.Println(sql)

	return count
}

func DeleteSymbolInfo(id int64) (count int64) {
	if id <= 0 {
		return 0
	}

	db := DB()
	query := db.Table("symbol_info")
	query.Where("id", id)
	count, _ = query.Delete()

	// output sql
	//sql, _, _ := query.BuildSql()
	//fmt.Println(sql)

	return count
}

func loadAllSymbolInfo() (symbolList []SymbolInfo) {
	db := DB()
	query := db.Table(&symbolList)
	query.Where("status", "=", config.STATUS_ENABLE)
	query.Select()
	return symbolList
}
