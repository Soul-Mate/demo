package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
)

type OrderController struct {
}

func (c *OrderController) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userId := request.URL.Query().Get("user_id")
		goodsId := request.URL.Query().Get("goods_id")
		writer.Header().Set("Content-Type", "application/json")
		if userId == "" || goodsId == "" {
			respData := map[string]interface{}{
				"status":  "false",
				"code":    200,
				"message": "缺少参数",
				"data":    map[string]interface{}{},
			}
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		id, _ := strconv.Atoi(userId)

	}
}
