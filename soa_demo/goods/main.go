package main

import (
	"net/http"
	"log"
	"github.com/Soul-Mate/demo/soa_demo/redis"
	"encoding/json"
)

func main() {
	registerService()
	goodsController := new(GoodsController)
	http.HandleFunc("/goods", goodsController.Index())
	http.HandleFunc("/goods/show", goodsController.Show())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerService() {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()
	reply, _ := redisConn.Do("HEXISTS", "soa_demo_services", "goods")
	if exists := reply.(int64); exists == 1 {
		return
	}
	data, _ := json.Marshal(map[string]string{
		"name": "goods_service",
		"addr": "localhost:8080",
		"host": "localhost",
		"port": "8080",
	})
	redisConn.Do("HMSET", "soa_demo_services", "goods", string(data))
}
