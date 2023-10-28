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
		res = common.SocketRouter(req, services.GetMissionsWithMachine)
	case req := <-c[mid]["setMissions"]:
		res = common.SocketRouter(req, services.SetMissionsWithMachine)

	case req := <-c[mid]["getTimers"]:
		res = common.SocketRouter(req, services.GetTimersWithMachine)
	case req := <-c[mid]["setTimers"]:
		res = common.SocketRouter(req, services.SetTimersWithMachine)
	//
	//case req := <-c[mid]["getMessages"]:
	//	res = common.SocketRouter(req, services.GetMessages)
	//case req := <-c[mid]["addMessages"]:
	//	res = common.SocketRouter(req, services.AddMessages)
	//case req := <-c[mid]["delMessages"]:
	//	res = common.SocketRouter(req, services.DelMessages)

	case req := <-c[mid]["getMachines"]:
		res = common.SocketRouter(req, services.GetMachines)
	case req := <-c[mid]["setMachines"]:
		res = common.SocketRouter(req, services.SetMachines)

	case req := <-c[mid]["upload1"]:
		res = common.SocketRouter(req, services.UploadPic1)
	case req := <-c[mid]["upload2"]:
		res = common.SocketRouter(req, services.UploadPic2)

	case req := <-c[mid]["questions"]:
		res = common.SocketRouter(req, services.UploadQuestion)

	case req := <-c[mid]["addNotes"]:
		res = common.SocketRouter(req, services.AddNotes)
	case req := <-c[mid]["getNotes"]:
		res = common.SocketRouter(req, services.GetNotes)

	case req := <-c[mid]["userRecord"]:
		res = common.SocketRouter(req, services.GetUserRecord)

	case req := <-c[mid]["send"]:
		res = common.SocketRouter(req, func(m map[string]string) (interface{}, error) {
			return m, nil
		})
	case req := <-c[mid]["timeout"]:
		res = common.SocketRouter(req, func(m map[string]string) (interface{}, error) {
			return nil, errors.New("timeout")
		})

	case req := <-c[mid]["eventError"]:
		res = common.SocketRouter(req, func(m map[string]string) (interface{}, error) {
			return nil, errors.New("event not exist")
		})
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		common.SendConn(conn, err.Error(), mid, 3)
	} else {
		common.SendConn(conn, string(resJson), mid, 4)
	}

}
