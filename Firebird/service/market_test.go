package service_test

import (
	"Firebird/config"
	"Firebird/service"
	"Firebird/utils"
	"Firebird/websocket"
	"fmt"
	"testing"
	"time"
)

//测试获取合约信息接口
func Test_FutureContractInfo(t *testing.T) {
	contractInfo := service.FutureContractInfo("BTC", "", "")
	fmt.Println("获取合约信息: ", contractInfo)

}

//测试获取合约指数信息接口
func Test_FutureContractIndex(t *testing.T) {
	//contract_index := services.FutureContractIndex("BTC")
	contract_index := service.FutureContractIndex("BTC")
	fmt.Println("获取合约指数信息: ", contract_index)

}

//获取订单明细信息
func Test_FutureContractOrderDetail(t *testing.T) {

	contract_order_detail := service.FutureContractOrderDetail("BTC", "123556", "1", "100", "1539345271124", "1")
	fmt.Println("获取订单明细信息: ", contract_order_detail)

}

//合约取消订单
func Test_FutureContractCancel(t *testing.T) {
	contract_cancel := service.FutureContractCancel("123456", "BTC", "123456")
	fmt.Println("合约取消订单: ", contract_cancel)

}

//合约全部撤单
func Test_FutureContractCancelall(t *testing.T) {
	contract_cancelall := service.FutureContractCancelall("BTC")
	fmt.Println("合约全部撤单: ", contract_cancelall)

}

//获取合约当前未成交委托
func Test_FutureContractOpenorders(t *testing.T) {
	contract_openorders := service.FutureContractOpenorders("BTC", "1", "100")
	fmt.Println("获取合约当前未成交委托: ", contract_openorders)

}

//获取合约历史委托
func Test_FutureContractHisorders(t *testing.T) {
	contract_hisorders := service.FutureContractHisorders("BTC", "0", "1", "0", "90", "1", "50")
	fmt.Println("获取合约历史委托: ", contract_hisorders)
	time.Sleep(time.Second)

}

//测试合约下单接口
func Test_FutureContractOrder(t *testing.T) {
	//合约下单
	contract_order := service.FutureContractOrder("BTC", "this_week", "BTC181214", "", "6188", "12",
		"buy", "open", "10", "limit")
	fmt.Println("合约下单: ", contract_order)

}

//测试批量下单接口
func Test_FutureContractBatchorder(t *testing.T) {
	//合约批量下单
	ordersData := make([]*service.Order, 0)
	order1 := &service.Order{

		Symbol:         "BTC",
		ContractType:   "quarter",
		ContractCode:   "BTC181228",
		ClientOrderId:  "10",
		Price:          "6188",
		Volume:         "1",
		Direction:      "buy",
		Offset:         "open",
		LeverRate:      "10",
		OrderPriceType: "limit",
	}

	ordersData = append(ordersData, order1)
	order2 := &service.Order{

		Symbol:         "BTC",
		ContractType:   "quarter",
		ContractCode:   "BTC181228",
		ClientOrderId:  "11",
		Price:          "6189",
		Volume:         "2",
		Direction:      "buy",
		Offset:         "open",
		LeverRate:      "10",
		OrderPriceType: "limit",
	}
	ordersData = append(ordersData, order2)
	fmt.Println("ordersData:", ordersData)

	contract_batchorder := service.FutureContractBatchorder(ordersData)
	fmt.Println("合约批量下单ordersDataResult: ", contract_batchorder)

}

// test get
func Test_FutureMarketHistoryKline(t *testing.T) {
	result := service.FutureMarketHistoryKline("eosusdt", "1day", 2)
	println(result)
}

//测试 WebSocket 行情,交易 API
func Test_Websocket(t *testing.T) {
	//websocket.WSRun()   //无需本地IP地址，直接运行
	ipList, _ := utils.GetLocalIPv4s()
	ip := config.Local_IP
	if len(ipList) > 0 {
		ip = ipList[0]
	}
	println(ip)
	websocket.WSRunWithIP(ip) //配置文件须填写本地IP地址，WS运行太久，外部原因可能断开，支持自动重连
}

//测试 WebSocket 订单推送 API
func Test_Websocket_order(t *testing.T) {
	websocket.WSWithOrder()
}
