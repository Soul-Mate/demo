package mian

import (
	"time"
	"github.com/Soul-Mate/demo/soa_demo/models"
	"github.com/Soul-Mate/demo/soa_demo/db"
)

type OrderModel struct {
	Id         int
	MemberId   int
	OrderNum   time.Duration
	PayStatus  int
	CreateTime string
}

func (o *OrderModel) Create(member *models.MemberModel) (*OrderModel, error){
	o.MemberId = member.Id
	o.OrderNum = time.Duration(time.Now().Unix())
	o.PayStatus = 0
	o.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec("INSERT INTO orders (member_id, order_num, pay_status, create_time) VALUES(?, ?, ?, ?)",
		o.MemberId, o.OrderNum, o.PayStatus, o.CreateTime)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return o, nil
}
