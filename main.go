package main

import (
	"socketAPI/app/myCron"
	"socketAPI/common"
	"socketAPI/config"
	"socketAPI/server"
)

func main() {

	c := make(map[string]map[string]chan map[string]string)

	common.InitDB()

	myCron.Start()

	go server.HttpConnect(config.MyConfig.NET.Http)

	server.Socket(config.MyConfig.NET.Socket, c)

}
