package service

import (
	"Firebird/config"
	"Firebird/logger"
	"Firebird/utils"
	"strconv"
)

var log = logger.NewLogger("[service]")

func GetHistoryKline(strSymbol, strPeriod string, nSize int) string {

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/market/history/kline"
	strUrl := config.GetMarketUrl() + strRequestUrl

	jsonKLineReturn := utils.HttpGetRequest(strUrl, mapParams)
	return jsonKLineReturn
}
