package web

import (
	"Firebird/config"
	"Firebird/db"
	"Firebird/utils"
	"github.com/gin-gonic/gin"
)

func listUserTrade(c *gin.Context) {
	query := db.UserTradeQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserId = utils.GetParamInt64(c, "userId")
	query.SymbolId = utils.GetParamInt64(c, "symbolId")
	query.ScheduleId = utils.GetParamInt64(c, "scheduleId")

	query.Type = utils.GetParamInt(c, "type")
	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	query.StartTime = utils.GetParamTime(c, "startTime")
	query.EndTime = utils.GetParamTime(c, "endTime")
	count, resultList := db.QueryUserTrade(&query)
	dataList := make([]db.UserTradeVO, 0)
	if count > 0 {
		for _, trade := range resultList {
			dataList = append(dataList, convertToTradeVO(&trade))
		}
	}

	c.JSON(200, JSONResult{
		"retCode":    0,
		"message":    "SUCCESS",
		"dataList":   dataList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func convertToTradeVO(trade *db.UserTrade) (tradeVO db.UserTradeVO) {
	tradeVO.Id = trade.Id
	tradeVO.GmtCreate = trade.GmtCreate
	tradeVO.GmtModified = trade.GmtModified
	tradeVO.Type = trade.Type
	tradeVO.ScheduleId = trade.ScheduleId
	tradeVO.UserId = trade.UserId
	tradeVO.SymbolId = trade.SymbolId
	tradeVO.Price = trade.Price
	tradeVO.Amount = trade.Amount
	tradeVO.HoldPrice = trade.HoldPrice
	tradeVO.HoldAmount = trade.HoldAmount
	tradeVO.Reason = trade.Reason
	tradeVO.Status = trade.Status

	symbolInfo := db.GetSymbolFromCacheById(trade.SymbolId)
	if symbolInfo.Id > 0 {
		tradeVO.SymbolName = symbolInfo.SymbolName
		tradeVO.SymbolDesc = symbolInfo.SymbolDesc
		tradeVO.SymbolIcon = symbolInfo.SymbolIcon
		tradeVO.SymbolGroup = symbolInfo.SymbolGroup
	}

	return tradeVO
}

func updateUserTrade(c *gin.Context) {
	userTrade := db.UserTrade{}

	userTrade.Id = utils.GetParamInt64(c, "id")
	if userTrade.Id <= 0 {
		c.JSON(200, JSONResult{
			"retCode": 1,
			"message": "参数错误",
		})
		return
	}

	userTrade.UserId = utils.GetParamInt64(c, "userId")
	userTrade.SymbolId = utils.GetParamInt64(c, "symbolId")
	userTrade.ScheduleId = utils.GetParamInt64(c, "scheduleId")

	userTrade.Status = utils.GetParamInt(c, "status")
	userTrade.Type = utils.GetParamInt(c, "type")

	userTrade.Price = utils.GetParamFloat64(c, "price")
	userTrade.Amount = utils.GetParamFloat64(c, "amount")

	result := db.UpdateUserTrade(&userTrade)
	if result > 0 {
		result = CODE_SUCCESS
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    userTrade,
	})
}

func addUserTrade(c *gin.Context) {
	userTrade := db.UserTrade{
		ScheduleId: 0,
	}

	userTrade.Id = utils.GetParamInt64(c, "id")
	userTrade.UserId = utils.GetParamInt64(c, "userId")
	userTrade.SymbolId = utils.GetParamInt64(c, "symbolId")

	userTrade.Status = utils.GetParamInt(c, "status")
	userTrade.Type = utils.GetParamInt(c, "type")

	userTrade.Price = utils.GetParamFloat64(c, "price")
	userTrade.Amount = utils.GetParamFloat64(c, "amount")

	// fetch user account
	account := db.GetUserAccountByUid(userTrade.UserId, userTrade.SymbolId)
	if account.Id > 0 {
		if userTrade.Type == config.TRADE_BUY {
			doBuyTrade(&userTrade, &account)
		} else if userTrade.Type == config.TRADE_SOLD {
			// hold amount check
			if userTrade.Amount > account.HoldAmount {
				c.JSON(200, JSONResult{
					"retCode": CODE_FAILED,
					"message": "余额不足",
				})
				return
			}
			doSoldTrade(&userTrade, &account)
		}
	}

	result := db.InsertUserTrade(&userTrade)
	if result > 0 {
		result = CODE_SUCCESS
		db.UpdateUserAccount(&account)
	} else {
		result = CODE_FAILED
	}
	c.JSON(200, JSONResult{
		"retCode": result,
		"message": "SUCCESS",
		"data":    "操作成功",
	})
}

func doBuyTrade(userTrade *db.UserTrade, account *db.UserAccount) {
	tradeTotal := userTrade.Amount * userTrade.Price

	account.HoldAmount += userTrade.Amount
	account.Total += tradeTotal
	account.HoldPrice = account.Total / account.HoldAmount

	userTrade.HoldAmount = account.HoldAmount
	userTrade.HoldPrice = account.HoldPrice
}

func doSoldTrade(userTrade *db.UserTrade, account *db.UserAccount) {
	amount := userTrade.Amount

	soldTotal := amount * userTrade.Price
	buyTotal := amount * account.HoldPrice

	account.TotalBenefit += soldTotal - buyTotal

	account.HoldAmount -= amount
	account.Total -= buyTotal
	if 0 == account.HoldAmount {
		account.HoldPrice = 0
	} else {
		account.HoldPrice = account.Total / account.HoldAmount
	}

	userTrade.HoldAmount = account.HoldAmount
	userTrade.HoldPrice = account.HoldPrice
}
