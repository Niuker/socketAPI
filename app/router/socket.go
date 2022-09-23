package router

import (
	"encoding/json"
	"net"
	"socketAPI/app/services"
	"socketAPI/app/structure"
	"socketAPI/common"
)

func RegisterSocketRoutes(conn net.Conn, mid string, c map[string]map[string]chan map[string]string) {

	res := structure.ResData{Data: make(map[string]string)}

	select {
	case req := <-c["xc"]["getMissions"]:

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

	resJson, err := json.Marshal(res)
	if err != nil {
		common.SendConn(conn, err.Error())
	} else {
		common.SendConn(conn, string(resJson))
	}

}
