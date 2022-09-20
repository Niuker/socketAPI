package router

import (
	"WebsocketDemo/common"
	"WebsocketDemo/services"
	"WebsocketDemo/structure"
	"encoding/json"
	"net"
)

func RegisterSocketRoutes(conn net.Conn, mid string, c map[string]map[string]chan map[string]string) {

	res := structure.ResData{Data: make(map[string]string)}

	//time.Sleep(time.Second * 5)

	common.Log(mid, 22, c)

	select {
	case req := <-c["xc"]["getMissions"]:
		common.Log(mid, 33)

		if data, err := services.GetMissions(req); err == nil {
			res.Data = data
		} else {
			res.Code = 1
			res.Error = err.Error()
		}
	case req := <-c[mid]["setMissions"]:
		if data, err := services.SetMissions(req); err == nil {
			res.Data = data
		} else {
			res.Code = 1
			res.Error = err.Error()
		}
		//default:
		//	res.Code = 1
		//	res.Error = "event not fount"
		//}
	}
	common.Log(mid, 11)

	resJson, err := json.Marshal(res)
	if err != nil {
		common.SendConn(conn, err.Error())
	} else {
		common.SendConn(conn, string(resJson))
	}

}
