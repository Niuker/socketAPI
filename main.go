package main

import (
	"socketAPI/app/myCron"
	"socketAPI/app/structure"
	"socketAPI/common"
	"socketAPI/config"
	"socketAPI/server"
)

func init() {
	common.InitDB()
	myCron.Start()
}

func main() {

	c := make(map[string]map[string]chan structure.ReqData)

	go server.HttpConnect(config.MyConfig.NET.Http)

	server.Socket(config.MyConfig.NET.Socket, c)

}
