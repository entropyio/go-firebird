package api

import (
	"Firebird/db"
	"Firebird/utils"
	"Firebird/web"
	"github.com/gin-gonic/gin"
)

func ListConfigInfo(c *gin.Context) {
	query := db.ConfigInfoQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.Ckey = utils.GetParamString(c, "ckey")
	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	count, dataList := db.QueryConfigInfo(&query)
	c.JSON(200, web.JSONResult{
		"retCode":    web.CODE_SUCCESS,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func SaveConfigInfo(c *gin.Context) {
	ConfigInfo := db.ConfigInfo{}

	ConfigInfo.Id = utils.GetParamInt64(c, "id")
	ConfigInfo.Ckey = utils.GetParamString(c, "ckey")
	ConfigInfo.Cvalue = utils.GetParamString(c, "cvalue")
	ConfigInfo.Status = utils.GetParamInt(c, "status")
	ConfigInfo.Operator = "system" // TODO: gin session
	var result int64 = 0
	if ConfigInfo.Id > 0 {
		result = db.UpdateConfigInfo(&ConfigInfo)
	} else {
		ConfigInfo.Creator = ConfigInfo.Operator
		result = db.InsertConfigInfo(&ConfigInfo)
	}

	if result > 0 {
		result = web.CODE_SUCCESS
	} else {
		result = web.CODE_FAILED
	}
	c.JSON(200, web.JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    ConfigInfo,
	})
}

func DeleteConfigInfo(c *gin.Context) {
	id := utils.GetParamInt64(c, "id")
	if id <= 0 {
		c.JSON(200, web.JSONResult{
			"retCode": web.CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	result := db.DeleteConfigInfo(id)
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
