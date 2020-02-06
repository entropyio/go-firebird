package service

import (
	"Firebird/config"
	"Firebird/db"
	"Firebird/utils"
	"github.com/robfig/cron"
	"strconv"
	"time"
)

func NotifySchedulePrice(symbol string, price float64) {
	//log.Debug("schedule notify", symbol, price)

}

/**
计算昨日收益定时任务
 */
func StartScheduleTask() {
	c := cron.New()

	// load cache - every 2 minutes
	spec1 := "0 */2 * * * ?"
	c.AddFunc(spec1, func() {
		log.Info("load all db data to cache")
		db.LoadAllToCache()
	})

	// 每天0：10 - update daily benefit
	spec2 := "0 10 0 /1 * ?"
	c.AddFunc(spec2, func() {
		log.Info("corn update yesterday benefit: ", utils.GetDateTimeStr(time.Now()))
		syncAllSymbolKline(config.SYMBOL_SYNC_COUNT)
		updateAllUserBenefit(1)
	})

	c.Start()

	defer c.Stop()

	select {}
}

func syncAllSymbolKline(number int) {
	query := db.SymbolInfoQuery{
		Status: 1,
	}
	count, symbolList := db.QuerySymbolInfo(&query)
	if count > 0 {
		for _, symbol := range symbolList {
			syncSymbolKline(symbol.SymbolName, number)
		}
	}
}

func updateAllUserBenefit(number int) {
	query := db.UserAccountQuery{
		Status: 1,
	}
	count, accountList := db.QueryUserAccount(&query)
	if count > 0 {
		now := time.Now()
		for i := number - 1; i > 0; i-- {
			duration, _ := time.ParseDuration(strconv.Itoa(-48*i) + "h")
			beforeTime := now.Add(duration)
			duration, _ = time.ParseDuration(strconv.Itoa(-24*i) + "h")
			yestTime := now.Add(duration)

			for _, account := range accountList {
				updateUserBenefit(&account, yestTime, beforeTime)
				userData := updateUserData(&account, yestTime, beforeTime)
				updateUserTotal(userData)
			}
		}

		// update total to db
		for _, totalData := range userTotalMap {
			totalData.HoldRate = (totalData.HoldAmount - totalData.HoldRate) * 100 / totalData.HoldRate
			totalData.HoldAmount = 0
			count := db.InsertUserData(totalData)
			if count == 0 {
				log.Error("Update user total error. time = " + utils.GetDateStr(now))
			} else {
				log.Info("Update user total success. time = " + utils.GetDateStr(now))
			}
		}
	}
}

func updateUserBenefit(account *db.UserAccount, now time.Time, base time.Time) {
	basePrice := GetSymbolPriceByTime(account.SymbolId, base)
	nowPrice := GetSymbolPriceByTime(account.SymbolId, now)
	account.YestBenefit = (nowPrice - basePrice) * account.HoldAmount
	count := db.UpdateUserAccount(account)
	if count == 0 {
		log.Error("Update user benefit error. time = " + utils.GetDateStr(now))
	} else {
		log.Info("Update user benefit success. time = " + utils.GetDateStr(now))
	}
}

func updateUserData(account *db.UserAccount, now time.Time, base time.Time) *db.UserData {
	klineData := GetSymbolKlineByTime(account.SymbolId, now)
	if klineData.Id > 0 {
		userData := db.UserData{
			GmtCreate:   time.Unix(klineData.Id, 0),
			GmtModified: time.Unix(klineData.Id, 0),
			UserId:      account.UserId,
			SymbolId:    account.SymbolId,
			OpenPrice:   klineData.Open,
			ClosePrice:  klineData.Close,
			HighPrice:   klineData.High,
			LowPrice:    klineData.Low,
			HoldPrice:   account.HoldPrice,
			HoldAmount:  account.HoldAmount,
			Status:      config.STATUS_ENABLE,
		}

		if account.HoldPrice-0.0001 > 0 && klineData.Close > 0 {
			userData.HoldRate = (klineData.Close - account.HoldPrice) * 100 / account.HoldPrice
			userData.HoldBenefit = (klineData.Close - account.HoldPrice) * account.HoldAmount
		}

		id := db.InsertUserData(&userData)
		userData.Id = id
		if id == 0 {
			log.Error("Update user data error. time = " + utils.GetDateStr(userData.GmtCreate))
		} else {
			log.Info("Update user data success. time = " + utils.GetDateStr(userData.GmtCreate))
		}
		return &userData
	}
	return &db.UserData{}
}

var userTotalMap = make(map[string]*db.UserData)

func updateUserTotal(userData *db.UserData) {
	if userData.Id > 0 {
		key := strconv.FormatInt(userData.UserId, 10) +
			"-" + utils.GetDateStr(userData.GmtCreate)
		totalData := userTotalMap[key]
		if nil == totalData {
			totalData = &db.UserData{
				GmtCreate:   userData.GmtCreate,
				GmtModified: userData.GmtModified,
				UserId:      userData.UserId,
				SymbolId:    0, // 0 for total
				HoldAmount:  0,
				HoldBenefit: 0,
				HoldRate:    0,
				Status:      config.STATUS_ENABLE,
			}
			userTotalMap[key] = totalData
		}

		totalData.HoldBenefit += userData.HoldBenefit
		totalData.HoldAmount += userData.ClosePrice * userData.HoldAmount // close total
		totalData.HoldRate += userData.HoldPrice * userData.HoldAmount    // hold total
	}
}
