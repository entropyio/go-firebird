package service

import (
	"Firebird/db"
	"time"
	"Firebird/utils"
	"Firebird/data"
	"strings"
	"encoding/json"
)

/**
current price
 */
var priceMap = make(map[string]float64)

func GetSymbolPrice(symbol string) (price float64) {
	return priceMap[symbol]
}

func GetSymbolPriceById(symbolId int64) (price float64) {
	symbolInfo := db.GetSymbolFromCacheById(symbolId)
	if symbolInfo.Id > 0 {
		return priceMap[symbolInfo.SymbolName]
	}
	return 0.0
}

func NotifySymbolPrice(symbol string, price float64) {
	priceMap[symbol] = price
}

/**
yesterday price
 */
var timePriceMap = make(map[string]data.KlineTickData)

func GetSymbolPriceByTime(symbolId int64, st time.Time) (price float64) {
	klineData := getSymbolKline(symbolId, st)
	if klineData.Id > 0 {
		return klineData.Close
	}
	return 0.0
}

func GetSymbolKlineByTime(symbolId int64, st time.Time) (data.KlineTickData) {
	return getSymbolKline(symbolId, st)
}

func getSymbolKline(symbolId int64, st time.Time) (data.KlineTickData) {
	symbolInfo := db.GetSymbolFromCacheById(symbolId)
	if symbolInfo.Id == 0 {
		return data.KlineTickData{}
	}

	key := symbolInfo.SymbolName + "-" + utils.GetDateStr(st)
	if _, ok := timePriceMap[key]; ok {
		return timePriceMap[key]
	}

	return timePriceMap[key]
}

func syncSymbolKline(symbolName string, count int) {
	if count <= 0 {
		count = 3
	}

	msg := GetHistoryKline(symbolName, "1day", count)
	//log.Info(msg)
	klineData := data.HistoryKlineData{}
	json.Unmarshal([]byte(msg), &klineData)
	dataList := klineData.Data
	nowStr := utils.GetDateStr(time.Now())

	if len(dataList) > 0 {
		symbols := strings.Split(klineData.Ch, ".")
		for _, item := range dataList {
			date := utils.GetDateStr(time.Unix(item.Id, 0))
			if strings.Compare(date, nowStr) == 0 {
				continue
			}

			tempKey := symbols[1]
			tempKey += "-"
			tempKey += date
			log.Info(tempKey, item.Close)
			timePriceMap[tempKey] = item
		}
	}
}
