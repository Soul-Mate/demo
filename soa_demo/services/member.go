package services

import (
	"github.com/Soul-Mate/demo/soa_demo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/Soul-Mate/demo/soa_demo/redis"
	"errors"
)

type MemberService struct {
}

func (m *MemberService) GetMember(memberId int) (*models.MemberModel, error) {
	// 查询注册的member服务
	redisConn := redis.Pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("HGET", "soa_demo_services", "member")
	if err != nil {
		return nil, err
	}
	serviceInfo := string(reply.([]uint8))
	if serviceInfo == "" {
		return nil, errors.New("请求繁忙,请稍后")
	}

	serviceInfoMap := make(map[string]interface{})
	json.Unmarshal(reply.([]uint8), &serviceInfoMap)
	url := fmt.Sprintf("http://%s/members/show?id=%d", serviceInfoMap["addr"], memberId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	respMap := make(map[string]interface{})
	json.Unmarshal(data, &respMap)

	// 不做重试处理
	if respMap["status"].(bool) != true {
		return nil, errors.New("请求繁忙,请稍后")
	}
	fmt.Println(respMap["data"])
	data, _ = json.Marshal(respMap["data"])
	memberM := new(models.MemberModel)
	json.Unmarshal(data, memberM)
	return memberM, nil
}
