package web

import (
	"Firebird/config"
	"Firebird/db"
	"Firebird/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func listUserSchedule(c *gin.Context) {
	query := db.UserScheduleQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserId = utils.GetParamInt64(c, "userId")
	query.SymbolId = utils.GetParamInt64(c, "symbolId")

	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	query.StartTime = utils.GetParamTime(c, "startTime")
	query.EndTime = utils.GetParamTime(c, "endTime")

	count, resultList := db.QueryUserSchedule(&query)
	dataList := make([]db.UserScheduleVO, 0)
	if count > 0 {
		for _, schedule := range resultList {
			dataList = append(dataList, convertToScheduleVO(&schedule))
		}
	}

	c.JSON(200, JSONResult{
		"retCode":    CODE_SUCCESS,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func convertToScheduleVO(schedule *db.UserSchedule) (scheduleVO db.UserScheduleVO) {
	scheduleVO.Id = schedule.Id
	scheduleVO.GmtCreate = schedule.GmtCreate
	scheduleVO.GmtModified = schedule.GmtModified
	scheduleVO.Name = schedule.Name
	scheduleVO.Type = schedule.Type
	scheduleVO.Amount = schedule.Amount
	scheduleVO.UserId = schedule.UserId
	scheduleVO.SymbolId = schedule.SymbolId
	scheduleVO.Status = schedule.Status

	symbolInfo := db.GetSymbolFromCacheById(schedule.SymbolId)
	if symbolInfo.Id > 0 {
		scheduleVO.SymbolName = symbolInfo.SymbolName
		scheduleVO.SymbolDesc = symbolInfo.SymbolDesc
		scheduleVO.SymbolIcon = symbolInfo.SymbolIcon
		scheduleVO.SymbolGroup = symbolInfo.SymbolGroup
	}

	return scheduleVO
}

func saveUserSchedule(c *gin.Context) {
	userSchedule := db.UserSchedule{}

	userSchedule.Id = utils.GetParamInt64(c, "id")
	userSchedule.UserId = utils.GetParamInt64(c, "userId")
	userSchedule.SymbolId = utils.GetParamInt64(c, "symbolId")
	userSchedule.Status = utils.GetParamInt(c, "status")
	userSchedule.Name = utils.GetParamString(c, "name")

	var result int64 = 0
	if userSchedule.Id > 0 {
		result = db.UpdateUserSchedule(&userSchedule)
	} else {
		result = db.InsertUserSchedule(&userSchedule)
	}

	if result > 0 {
		result = CODE_SUCCESS
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    userSchedule,
	})
}

func deleteUserSchedule(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if id <= 0 {
		c.JSON(200, JSONResult{
			"retCode": CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	result := db.DeleteUserSchedule(id)
	if result > 0 {
		result = CODE_SUCCESS
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    result,
	})
}

func detailUserSchedule(c *gin.Context) {
	id := utils.GetParamInt64(c, "id")
	if id <= 0 {
		c.JSON(200, JSONResult{
			"retCode": CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	scheduleVO := getUserScheduleDetail(id)
	c.JSON(200, JSONResult{
		"retCode": CODE_SUCCESS,
		"message": "SUCCESS",
		"data":    scheduleVO,
	})
}

func getUserScheduleDetail(id int64) db.UserScheduleVO {
	// query schedule
	query := db.UserScheduleQuery{
		Id:     id,
		Status: config.STATUS_ENABLE,
	}
	count, resultList := db.QueryUserSchedule(&query)
	if count <= 0 {
		return db.UserScheduleVO{}
	}

	scheduleVO := convertToScheduleVO(&resultList[0])

	// query rules
	ruleQuery := db.RuleItemQuery{
		ScheduleId: scheduleVO.Id,
		UserId:     scheduleVO.UserId,
		SymbolId:   scheduleVO.SymbolId,
		Status:     config.STATUS_ENABLE,
	}
	ruleCount, ruleList := db.QueryRuleItem(&ruleQuery)
	if ruleCount > 0 {
		scheduleVO.Rules = ruleList
	}
	return scheduleVO
}

func listRuleItem(c *gin.Context) {
	query := db.RuleItemQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserId = utils.GetParamInt64(c, "userId")
	query.SymbolId = utils.GetParamInt64(c, "symbolId")
	query.ScheduleId = utils.GetParamInt64(c, "scheduleId")

	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	count, dataList := db.QueryRuleItem(&query)
	c.JSON(200, JSONResult{
		"retCode":    CODE_SUCCESS,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func saveRuleItem(c *gin.Context) {
	rule := db.RuleItem{}

	rule.Id = utils.GetParamInt64(c, "id")
	rule.UserId = utils.GetParamInt64(c, "userId")
	rule.SymbolId = utils.GetParamInt64(c, "symbolId")
	rule.ScheduleId = utils.GetParamInt64(c, "scheduleId")

	rule.RuleType = utils.GetParamInt(c, "ruleType")
	rule.JoinType = utils.GetParamInt(c, "joinType")
	rule.OpType = utils.GetParamInt(c, "opType")
	rule.Value = utils.GetParamString(c, "value")

	var result int64 = 0
	if rule.Id > 0 {
		result = db.UpdateRuleItem(&rule)
	} else {
		result = db.InsertRuleItem(&rule)
	}

	if result > 0 {
		result = CODE_SUCCESS
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    rule,
	})
}

func deleteRuleItem(c *gin.Context) {
	id := utils.GetParamInt64(c, "id")
	if id <= 0 {
		c.JSON(200, JSONResult{
			"retCode": CODE_FAILED,
			"message": "参数错误",
		})
		return
	}

	result := db.DeleteRuleItem(id)
	if result > 0 {
		result = CODE_SUCCESS
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    result,
	})
}
