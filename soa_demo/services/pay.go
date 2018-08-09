package services

import (
	"time"
	"github.com/Soul-Mate/demo/soa_demo/redis"
	"log"
	"strconv"
	"github.com/Soul-Mate/demo/soa_demo/models"
)

type PayService struct {
}

func (p *PayService) Pay(orderId int) error{
	orderM := new(models.OrderModel)
	_, err := orderM.Find(orderId)
	if err != nil {
		return err
	}
	// TODO 扣减用于账户余额
	orderM.PayStatus = 1
	orderM.Save()
	return nil
}

func (p *PayService) PayCron() {
	// 每1s执行一次
	t := time.NewTimer(time.Second)
	conn := redis.Pool.Get()
	defer conn.Close()
	for {
		select {
		case <-t.C:
			reply, err := conn.Do("LPOP", "pay_list")
			if err != nil {
				log.Println(err)
			} else if reply != nil {
				orderId, _ := strconv.Atoi(string(reply.([]uint8)))
				// 调用支付服务
				if err = p.Pay(orderId); err ==nil {
					// TODO 调用订单的回调接口,通知该笔订单处理完成
				} else {
					log.Println(err)
				}
			}
			// 重置计数器
			t = time.NewTimer(time.Second)
		}
	}
}
