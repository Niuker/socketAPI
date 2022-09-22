package main

import (
	"WebsocketDemo/common"
	"WebsocketDemo/config"
	"WebsocketDemo/server"
)

func main() {

	c := make(map[string]map[string]chan map[string]string)

	common.InitDB()

	go server.HttpConnect(config.MyConfig.NET.Http)

	server.Socket(config.MyConfig.NET.Socket, c)

}
