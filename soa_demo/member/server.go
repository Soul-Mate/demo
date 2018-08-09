package main

import (
	"github.com/Soul-Mate/demo/soa_demo/controllers"
	"net/http"
	"log"
	"encoding/json"
	"github.com/Soul-Mate/demo/soa_demo/redis"
)

func main() {
	registerService()
	memberC := new(controllers.MemberController)
	http.HandleFunc("/members/show", memberC.Show())
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func registerService() {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()
	data, _ := json.Marshal(map[string]string{
		"name": "goods_service",
		"addr": "localhost:8001",
		"host": "localhost",
		"port": "8001",
	})
	redisConn.Do("HMSET", "soa_demo_services", "member", string(data))
}
