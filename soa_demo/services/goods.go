package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/Soul-Mate/demo/soa_demo/models"
	"github.com/Soul-Mate/demo/soa_demo/redis"
	"errors"
)

type GoodsService struct {
}

func (g *GoodsService) GetGoods(goodsId int) (*models.GoodsModel, error) {
	redisConn := redis.Pool.Get()
	defer redisConn.Close()
	// 查询注册的goods服务
	reply, err := redisConn.Do("HGET", "soa_demo_services", "goods")
	if err != nil {
		return nil, err
	}
	serviceInfo := string(reply.([]uint8))
	if serviceInfo == "" {
		return nil, errors.New("请求失败,请稍后")
	}
	// 进行服务同步调用
	serviceInfoMap := make(map[string]interface{})
	json.Unmarshal(reply.([]uint8), &serviceInfoMap)
	url := fmt.Sprintf("http://%s/goods/show?id=%d", serviceInfoMap["addr"], goodsId)
	// 通过http api 查询商品信息
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	respMap := make(map[string]interface{})
	json.Unmarshal(data, &respMap)
	// 不做重试处理
	if respMap["status"].(bool) != true {

		return nil, errors.New("请求失败,请稍后")
	}
	data, _ = json.Marshal(respMap["data"])
	goodsM := new(models.GoodsModel)
	json.Unmarshal(data, goodsM)
	return goodsM, nil
}
