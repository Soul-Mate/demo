package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/Soul-Mate/demo/soa_demo/models"
)

type MemberController struct{}

func (m *MemberController) Show() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		memberIdPram := request.URL.Query().Get("id")
		writer.Header().Set("Content-Type", "application/json")
		respData := map[string]interface{}{
			"status":  false,
			"code":    200,
			"message": "",
			"data":    map[string]interface{}{},
		}
		if memberIdPram == "" {
			respData["message"] = "缺少参数"
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		memberId, _ := strconv.Atoi(memberIdPram)
		memberM := new(models.MemberModel)
		_, err := memberM.Find(memberId)
		if err != nil {
			respData["message"] = err.Error()
			data, _ := json.Marshal(respData)
			writer.Write(data)
			fmt.Fprint(writer)
			return
		}
		respData["message"] = "success."
		respData["status"] = true
		respData["data"] = map[string]interface{}{
			"member_id":     memberM.Id,
			"member_name":   memberM.Name,
			"member_level":  memberM.Level,
			"member_amount": memberM.Amount,
		}
		data, _ := json.Marshal(respData)
		writer.Write(data)
		fmt.Fprint(writer)
		return
	}
}
