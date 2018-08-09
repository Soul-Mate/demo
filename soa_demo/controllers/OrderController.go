package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/Soul-Mate/demo/soa_demo/models"
	"github.com/Soul-Mate/demo/soa_demo/services"
)

type OrderController struct {
}

func (c *OrderController) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userIdParam := request.URL.Query().Get("user_id")
		goodsIdParam := request.URL.Query().Get("goods_id")
		writer.Header().Set("Content-Type", "application/json")
		respData := map[string]interface{}{
			"status":  false,
			"code":    200,
			"message": "",
			"data":    map[string]interface{}{},
		}
		if userIdParam == "" || goodsIdParam == "" {
			respData["message"] = "缺少参数"
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		memberId, _ := strconv.Atoi(userIdParam)
		goodsId, _ := strconv.Atoi(goodsIdParam)
		// 查询注册的goods服务
		goodsService := new(services.GoodsService)
		goodsM, err := goodsService.GetGoods(goodsId)
		if err != nil {
			respData["message"] = err.Error()
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}

		// 查询注册的goods服务
		memberService := new(services.MemberService)
		memberM, err := memberService.GetMember(memberId)
		if err != nil {
			respData["message"] = err.Error()
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}

		// 创建订单
		orderM := new(models.OrderModel)
		_, err = orderM.Create(memberM, goodsM)
		if err != nil {
			respData["message"] = err.Error()
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		respData["message"] = "success."
		respData["status"] = "true"
		respData["data"] = map[string]interface{}{
			"member_id":   orderM.MemberId,
			"goods_id":    orderM.GoodsId,
			"price":       orderM.Price,
			"pay_status":  orderM.PayStatus,
			"order_num":   orderM.OrderNum,
			"create_time": orderM.CreateTime,
		}
		data, _ := json.Marshal(respData)
		writer.Write(data)
		fmt.Fprint(writer)
	}
}
