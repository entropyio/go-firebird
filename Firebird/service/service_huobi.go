package service

import (
	"strconv"
	"Firebird/config"
	"Firebird/utils"
	"Firebird/logger"
)

var log = logger.NewLogger("[service]")

func GetHistoryKline(strSymbol, strPeriod string, nSize int) string {

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/market/history/kline"
	strUrl := config.MARKET_URL + strRequestUrl

	jsonKLineReturn := utils.HttpGetRequest(strUrl, mapParams)
	return jsonKLineReturn
}
