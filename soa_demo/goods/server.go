package main

import (
	"net/http"
	"log"
	"github.com/Soul-Mate/demo/soa_demo/redis"
	"encoding/json"
	"github.com/Soul-Mate/demo/soa_demo/controllers"
)

func main() {
	registerService()
	goodsController := new(controllers.GoodsController)
	http.HandleFunc("/goods", goodsController.Index())
	http.HandleFunc("/goods/show", goodsController.Show())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func registerService() {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()
	data, _ := json.Marshal(map[string]string{
		"name": "goods_service",
		"addr": "localhost:8000",
		"host": "localhost",
		"port": "8000",
	})
	redisConn.Do("HMSET", "soa_demo_services", "goods", string(data))
}
