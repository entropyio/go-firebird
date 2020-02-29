package api

import (
	"Firebird/db"
	"Firebird/utils"
	"Firebird/web"
	"github.com/gin-gonic/gin"
)

func ListUserData(c *gin.Context) {
	query := db.UserDataQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserId = utils.GetParamInt64(c, "userId")
	query.SymbolId = utils.GetParamInt64(c, "symbolId")

	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	query.StartTime = utils.GetParamTime(c, "startTime")
	query.EndTime = utils.GetParamTime(c, "endTime")

	count, dataList := db.QueryUserData(&query)

	c.JSON(200, web.JSONResult{
		"retCode":    0,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}
