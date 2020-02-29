package api

import (
	"Firebird/db"
	"Firebird/utils"
	"Firebird/web"
	"github.com/gin-gonic/gin"
)

func ListUserInfo(c *gin.Context) {
	query := db.UserInfoQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserName = utils.GetParamString(c, "userName")
	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	count, dataList := db.QueryUserInfo(&query)
	c.JSON(200, web.JSONResult{
		"retCode":    web.CODE_SUCCESS,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func SaveUserInfo(c *gin.Context) {
	UserInfo := db.UserInfo{}

	UserInfo.Id = utils.GetParamInt64(c, "id")
	UserInfo.UserName = utils.GetParamString(c, "userName")
	UserInfo.UserDesc = utils.GetParamString(c, "userDesc")
	UserInfo.Status = utils.GetParamInt(c, "status")

	var result int64 = 0
	if UserInfo.Id > 0 {
		result = db.UpdateUserInfo(&UserInfo)
	} else {
		result = db.InsertUserInfo(&UserInfo)
	}

	if result > 0 {
		result = web.CODE_SUCCESS
	} else {
		result = web.CODE_FAILED
	}
	c.JSON(200, web.JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    UserInfo,
	})
}

func DeleteUserInfo(c *gin.Context) {
	id := utils.GetParamInt64(c, "id")
	if id <= 0 {
		c.JSON(200, web.JSONResult{
			"retCode": web.CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	result := db.DeleteUserInfo(id)
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
