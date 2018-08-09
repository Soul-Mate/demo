package main

import (
	"github.com/Soul-Mate/demo/soa_demo/services"
	"time"
)

func main() {
	payService := new(services.PayService)
	go payService.PayCron()
	time.Sleep(time.Second * 100)
}
