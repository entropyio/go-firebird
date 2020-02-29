package server

import (
	"Firebird/db"
	"Firebird/web/api"
	"fmt"
	"testing"
)

func TestStartHttpServer(t *testing.T) {
	db.LoadAllToCache()
	StartHttpServer("../../admin")
}

func TestDetailUserSchedule(t *testing.T) {
	vo := api.GetUserScheduleDetail(1)
	fmt.Println(vo)
}
