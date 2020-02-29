package db

import (
	"Firebird/logger"
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
)

var log = logger.NewLogger("[db]")

var err error
var engine *gorose.Engin

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	// mysql Dsn示例 "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"
	engine, err = gorose.Open(&gorose.Config{
		Driver: "sqlite3",
		Dsn:    "D:/blockchain/Firebird/go-firebird/Firebird/data/firebird.db",
	})
	if nil != err {
		panic(err)
	}
}

func DB() gorose.IOrm {
	return engine.NewOrm()
}
