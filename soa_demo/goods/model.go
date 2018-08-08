package main

import (
	"github.com/Soul-Mate/demo/soa_demo/db"
	"log"
	)

type GoodsModel struct {
	Id    int
	Name  string
	Price float64
	Stock int
}

func (g *GoodsModel) All() ([]*GoodsModel, error) {
	var allGoods []*GoodsModel
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM goods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		goodsM := new(GoodsModel)
		if err = rows.Scan(&goodsM.Id, &goodsM.Name, &goodsM.Price, &goodsM.Stock); err != nil {
			log.Printf("query goods error: %s", err.Error())
		} else {
			allGoods = append(allGoods, goodsM)
		}
	}
	return allGoods, nil
}

func (g *GoodsModel) Find(id int) (*GoodsModel, error) {
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow("SELECT * FROM goods where id = ?", id)
	if err = row.Scan(&g.Id, &g.Name, &g.Price, &g.Stock); err != nil {
		log.Printf("query goods error: %s", err.Error())
		return nil, err
	}
	return g, nil
}

