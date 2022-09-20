package main

import (
	"WebsocketDemo/client/model"
	"fmt"
	"time"
)

func main() {
	for true {
		var client = &model.SocketClient{
			Network: "tcp4",
			//Address: "47.94.99.64:8000",
			Address: "localhost:8000",
			//Address: ":9800",
			OnMessage: func(msg string) {
				fmt.Println("接收到服务端的消息", msg)
			},
		}
		err := client.Start()
		if err != nil {
			fmt.Println("连接服务端失败", err)
			return
		}

		fmt.Println("连接服务端成功", client)

		//for true {
		client.Connect.Write([]byte("{\"event\":\"getMissions\",\"params\":{\"id\":\"Rn+lqjJcpT0=\"},\"timestamp\":123}\n"))
		time.Sleep(time.Second * 1)

		client.Connect.Write([]byte("{\"event\":\"getMissions\",\"params\":{\"id\":\"Rn+lqjJcpT0=\"},\"timestamp\":123311113}\n"))
		//time.Sleep(time.Millisecond * 10)
		//
		client.Connect.Write([]byte("{\"Event\":\"setMissions\",\"params\":{\"id\":\"Rn+lqjJcpT0=\"},\"timestamp\":123333333}\n"))
		client.Connect.Write([]byte("{\"Event\":\"setMissions\",\"params\":{\"id\":\"Rn+lqjJcpT0=\",\"a\":\"1233\"},\"timestamp\":1233333555}\n"))

		//}
		client.Wg.Wait()
		fmt.Println("与服务器连接断开，3秒后重试")
		time.Sleep(time.Second * 3)
	}
}