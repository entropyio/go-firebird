package web

import (
	"fmt"
	"testing"
)

func TestStartHttpServer(t *testing.T) {
	StartHttpServer()
}

func TestDetailUserSchedule(t *testing.T) {
	vo := getUserScheduleDetail(1)
	fmt.Println(vo)
}
