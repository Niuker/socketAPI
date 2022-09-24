package router

import (
	"encoding/json"
	"errors"
	"net"
	"socketAPI/app/services"
	"socketAPI/app/structure"
	"socketAPI/common"
)

func RegisterSocketRoutes(conn net.Conn, mid string, c map[string]map[string]chan structure.ReqData) {
	var res structure.ResData
	select {
	case req := <-c[mid]["getMissions"]:
		res = common.SocketRouter(req, services.GetMissions)
	case req := <-c[mid]["setMissions"]:
		res = common.SocketRouter(req, services.SetMissions)

	case req := <-c[mid]["getTimers"]:
		res = common.SocketRouter(req, services.GetTimers)
	case req := <-c[mid]["setTimers"]:
		res = common.SocketRouter(req, services.SetTimers)

	case req := <-c[mid]["getMessages"]:
		res = common.SocketRouter(req, services.GetMessages)
	case req := <-c[mid]["addMessages"]:
		res = common.SocketRouter(req, services.AddMessages)
	case req := <-c[mid]["delMessages"]:
		res = common.SocketRouter(req, services.DelMessages)

	case req := <-c[mid]["getMachines"]:
		res = common.SocketRouter(req, services.GetMachines)
	case req := <-c[mid]["setMachines"]:
		res = common.SocketRouter(req, services.SetMachines)

	case req := <-c[mid]["send"]:
		res = common.SocketRouter(req, func(m map[string]string) (interface{}, error) {
			return m, nil
		})
	case req := <-c[mid]["timeout"]:
		res = common.SocketRouter(req, func(m map[string]string) (interface{}, error) {
			return nil, errors.New("timeout")
		})

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
