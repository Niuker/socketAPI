package main

import (
	"net"
	"socketAPI/app/myCron"
	"socketAPI/app/structure"
	"socketAPI/common"
	"socketAPI/config"
	"socketAPI/server"
)

func init() {
	common.InitDB()
	myCron.Start()
	//services.Doit()
}

func main() {
	c := make(map[string]map[string]chan structure.ReqData)
	connMap := make(map[string]net.Conn)

	go server.HttpConnect(config.MyConfig.NET.Http)

	server.Socket(config.MyConfig.NET.Socket, c, connMap)

}
