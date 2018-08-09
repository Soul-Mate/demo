package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
)

type PayController struct {
}

func (p *PayController) Pay() http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		orderIdPram := request.URL.Query().Get("id")
		writer.Header().Set("Content-Type", "application/json")
		respData := map[string]interface{}{
			"status":  false,
			"code":    200,
			"message": "",
			"data":    map[string]interface{}{},
		}
		if orderIdPram == "" {
			respData["message"] = "缺少参数"
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		orderId, _ := strconv.Atoi(orderIdPram)
		// 将支付请求打入队列, 在并发情况下可以消峰

	}
}
