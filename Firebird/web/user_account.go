package web

import (
	"Firebird/db"
	"Firebird/service"
	"Firebird/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func listUserAccount(c *gin.Context) {
	query := db.UserAccountQuery{}

	query.Id = utils.GetParamInt64(c, "id")
	query.UserId = utils.GetParamInt64(c, "userId")
	query.SymbolId = utils.GetParamInt64(c, "symbolId")

	query.Status = utils.GetParamInt(c, "status")
	query.PageNumber = utils.GetParamInt(c, "pageNumber")
	query.PageSize = utils.GetParamInt(c, "pageSize")

	query.StartTime = utils.GetParamTime(c, "startTime")
	query.EndTime = utils.GetParamTime(c, "endTime")

	count, resultList := db.QueryUserAccount(&query)
	c.JSON(200, JSONResult{
		"retCode":    0,
		"message":    "SUCCESS",
		"dataList":   resultList,
		"totalCount": count,
		"pageNum":    query.PageNumber,
	})
}

func getUserAccountByUid(c *gin.Context) {
	userId := utils.GetParamInt64(c, "userId")
	symbolId := utils.GetParamInt64(c, "symbolId")
	if userId <= 0 || symbolId <= 0 {
		c.JSON(200, JSONResult{
			"retCode": 1,
			"message": "参数错误",
		})
		return
	}

	record := db.GetUserAccountByUid(userId, symbolId)
	c.JSON(200, JSONResult{
		"retCode": 0,
		"message": "SUCCESS",
		"data":    record,
	})
}

func getCurrentAccount(c *gin.Context) {
	userId := utils.GetParamInt64(c, "userId")
	symbolId := utils.GetParamInt64(c, "symbolId")
	if userId <= 0 || symbolId <= 0 {
		c.JSON(200, JSONResult{
			"retCode": 1,
			"message": "参数错误",
		})
		return
	}

	account := db.GetUserAccountByUid(userId, symbolId)
	accountVO := calculateCurrentAccount(&account)

	c.JSON(200, JSONResult{
		"retCode": 0,
		"message": "SUCCESS",
		"data":    accountVO,
	})
}

func calculateCurrentAccount(account *db.UserAccount) (accountVO db.UserAccountVO) {
	// base info
	accountVO.YestBenefit = account.YestBenefit
	accountVO.TotalBenefit = account.TotalBenefit
	accountVO.GmtModified = account.GmtModified
	accountVO.GmtCreate = account.GmtCreate
	accountVO.HoldAmount = account.HoldAmount
	accountVO.HoldPrice = account.HoldPrice
	accountVO.SymbolId = account.SymbolId
	accountVO.UserId = account.UserId
	accountVO.Amount = account.Amount
	accountVO.Id = account.Id

	// symbol info
	symbolInfo := db.GetSymbolFromCacheById(account.SymbolId)
	if symbolInfo.Id > 0 {
		accountVO.SymbolName = symbolInfo.SymbolName
		accountVO.SymbolDesc = symbolInfo.SymbolDesc
		accountVO.SymbolIcon = symbolInfo.SymbolIcon
		accountVO.SymbolGroup = symbolInfo.SymbolGroup
	}

	price := service.GetSymbolPriceById(account.SymbolId)
	if price > 0 {
		account.Price = price
		account.Total = account.Price * account.HoldAmount
		account.Rate = (price - account.HoldPrice) * 100 / account.HoldPrice
		account.Benefit = (price - account.HoldPrice) * account.HoldAmount

		accountVO.Price = account.Price
		accountVO.Total = account.Total
		accountVO.Rate = account.Rate
		accountVO.Benefit = account.Benefit
	}
	return accountVO
}

func listCurrentAccounts(c *gin.Context) {
	userId := utils.GetParamInt64(c, "userId")
	if userId <= 0 {
		c.JSON(200, JSONResult{
			"retCode": 1,
			"message": "参数错误",
		})
		return
	}

	totalVO := db.UserAccountVO{
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}

	query := db.UserAccountQuery{
		UserId: userId,
	}

	dataList := make([]db.UserAccountVO, 0)
	count, accounts := db.QueryUserAccount(&query)
	if count > 0 {
		for index, account := range accounts {
			dataList = append(dataList, calculateCurrentAccount(&account))
			accounts[index] = account

			totalVO.Total += account.Total
			totalVO.Rate += account.HoldPrice * account.HoldAmount
			totalVO.TotalBenefit += account.TotalBenefit
			totalVO.YestBenefit += account.YestBenefit
			totalVO.Benefit += account.Benefit
		}
	}

	if 0 != totalVO.Rate {
		totalVO.Rate = (totalVO.Total - totalVO.Rate) * 100 / totalVO.Rate
	}

	c.JSON(200, JSONResult{
		"retCode":  0,
		"message":  "SUCCESS",
		"data":     totalVO,
		"dataList": dataList,
	})
}
