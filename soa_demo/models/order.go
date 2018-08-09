package models

import (
	"time"
	"github.com/Soul-Mate/demo/soa_demo/db"
	)

type OrderModel struct {
	Id         int
	MemberId   int
	GoodsId    int
	Price      float64
	OrderNum   time.Duration
	PayStatus  int
	CreateTime string
}

func (o *OrderModel) Create(member *MemberModel, goods *GoodsModel) (*OrderModel, error) {
	o.MemberId = member.Id
	o.OrderNum = time.Duration(time.Now().Unix())
	o.PayStatus = 0
	o.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o.GoodsId = goods.Id
	o.Price = goods.Price
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	sql := `INSERT INTO orders (member_id, goods_id, price, order_num, pay_status, create_time) VALUES(?, ?, ?, ?, ?, ?)`
	_, err = conn.Exec(sql, o.MemberId, o.GoodsId, o.Price, o.OrderNum, o.PayStatus, o.CreateTime)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return o, nil
}

func (o *OrderModel) Find(id int) (*OrderModel, error) {
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	row := conn.QueryRow("select id, member_id, goods_id, price, pay_status, create_time from orders where id = ?", id)
	err = row.Scan(&o.Id, &o.MemberId, &o.GoodsId, &o.Price, &o.PayStatus, &o.CreateTime)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OrderModel) Save() (*OrderModel, error){
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec("update orders set member_id = ?, goods_id = ?, price = ?, order_num = ?, pay_status = ?, create_time = ?",
		o.MemberId, o.GoodsId, o.Price, o.OrderNum, o.PayStatus, o.CreateTime)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return o, nil
}