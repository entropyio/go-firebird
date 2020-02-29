package api

import (
	"Firebird/db"
	"Firebird/utils"
	"Firebird/web"
	"github.com/gin-gonic/gin"
)

func ListSymbolInfo(c *gin.Context) {
	query := db.SymbolInfoQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.SymbolName = utils.GetParamString(c, "symbolName")
	query.SymbolGroup = utils.GetParamString(c, "symbolGroup")
	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	count, dataList := db.QuerySymbolInfo(&query)
	c.JSON(200, web.JSONResult{
		"retCode":    web.CODE_SUCCESS,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func SaveSymbolInfo(c *gin.Context) {
	symbolInfo := db.SymbolInfo{}

	symbolInfo.Id = utils.GetParamInt64(c, "id")
	symbolInfo.SymbolName = utils.GetParamString(c, "symbolName")
	symbolInfo.SymbolIcon = utils.GetParamString(c, "symbolIcon")
	symbolInfo.SymbolDesc = utils.GetParamString(c, "symbolDesc")
	symbolInfo.SymbolGroup = utils.GetParamString(c, "symbolGroup")
	symbolInfo.Status = utils.GetParamInt(c, "status")

	var result int64 = 0
	if symbolInfo.Id > 0 {
		result = db.UpdateSymbolInfo(&symbolInfo)
	} else {
		result = db.InsertSymbolInfo(&symbolInfo)
	}

	if result > 0 {
		result = web.CODE_SUCCESS
	} else {
		result = web.CODE_FAILED
	}
	c.JSON(200, web.JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    symbolInfo,
	})
}

func DeleteSymbolInfo(c *gin.Context) {
	id := utils.GetParamInt64(c, "id")
	if id <= 0 {
		c.JSON(200, web.JSONResult{
			"retCode": web.CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	result := db.DeleteSymbolInfo(id)
	if result > 0 {
		result = web.CODE_SUCCESS
	} else {
		result = web.CODE_FAILED
	}
	c.JSON(200, web.JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    result,
	})
}
