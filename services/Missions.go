package services

import "WebsocketDemo/common"

func GetMissions(req map[string]string) (interface{}, error) {
	common.Log(req, 7777)

	m := []int{1, 12, 345}
	return m, nil
}

func SetMissions(req map[string]string) (interface{}, error) {
	common.Log(req, 999999)
	m := map[string]string{"1231": "3213", "123das": "123"}
	return m, nil
}
