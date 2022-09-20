package main

import (
	"WebsocketDemo/config"
	"WebsocketDemo/server"
	"fmt"
)

func main() {

	c := make(map[string]map[string]chan map[string]string)

	fmt.Println(config.MyConfig.NET.Http)
	go server.HttpConnect(config.MyConfig.NET.Http)

	server.Socket(config.MyConfig.NET.Socket, c)

}
