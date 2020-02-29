package main

import (
	"Firebird/config"
	"Firebird/db"
	"Firebird/service"
	"Firebird/utils"
	"Firebird/web/server"
	"Firebird/websocket"
	"fmt"
)

func main() {
	fmt.Println("start Firebird......")

	//配置文件须填写本地IP地址，WS运行太久，外部原因可能断开，支持自动重连
	ipList, _ := utils.GetLocalIPv4s()
	ip := config.Local_IP
	if len(ipList) > 0 {
		ip = ipList[0]
	}
	fmt.Println("local ip:", ip)

	go websocket.WSRunWithIP(ip)
	go service.StartScheduleTask()

	db.LoadAllToCache()
	server.StartHttpServer("./admin")

	fmt.Println("Firebird started.")
}
