package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/Soul-Mate/demo/soa_demo/models"
)

type GoodsController struct {
}

func (c *GoodsController) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		goodsM := new(models.GoodsModel)
		allGoods, err := goodsM.All()
		if err != nil {
			respData := map[string]interface{}{
				"code":    200,
				"status":  "false",
				"message": err.Error(),
				"data": map[string]interface{}{
				},
			}
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		allGoodsMapSlice := make([]map[string]interface{}, 0)
		for _, goods := range allGoods {
			allGoodsMapSlice = append(allGoodsMapSlice, map[string]interface{}{
				"goods_id":    goods.Id,
				"goods_name":  goods.Name,
				"goods_price": goods.Price,
				"goods_stock": goods.Stock,
			})
		}
		respData := map[string]interface{}{
			"status":  200,
			"message": "success",
			"data":    allGoodsMapSlice,
		}
		data, _ := json.Marshal(respData)
		writer.Write(data)
		fmt.Fprint(writer)
	}

}

func (c *GoodsController) Show() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 由于demo, 不进行错误处理,实际项目要严格进行错误处理
		writer.Header().Set("Content-Type", "application/json")
		v := request.URL.Query().Get("id")
		if v == "" {
			fmt.Fprintf(writer, "缺少参数")
			return
		}
		goodsM := new(models.GoodsModel)
		id, _ := strconv.Atoi(v)
		_, err := goodsM.Find(id)
		if err != nil {
			respData := map[string]interface{}{
				"code":    200,
				"status":  false,
				"message": err.Error(),
				"data": map[string]interface{}{
				},
			}
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		respData := map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "success",
			"data": map[string]interface{}{
				"goods_id":    goodsM.Id,
				"goods_name":  goodsM.Name,
				"goods_price": goodsM.Price,
				"goods_stock": goodsM.Stock,
			},
		}
		data, _ := json.Marshal(respData)
		writer.Write(data)
		return
	}
}
