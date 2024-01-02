package main

import (
	"fmt"
	"socketAPI/client/model"
	"time"
)

func main() {
	//for true {
	var client = &model.SocketClient{
		Network: "tcp4",
		//Address: "47.94.99.64:8000",
		Address: "localhost:8000",
		//Address: ":9800",
		OnMessage: func(msg string) {
			fmt.Println("接收到服务端的消息", msg)
		},
	}
	//if err != nil {
	//	fmt.Println("连接服务端失败", err)
	//	return
	//}

	fmt.Println("连接服务端成功", client)
	//for true {
	err := client.Start()

	response, err := client.SendMessageAndWait("{\"event\":\"setMachines\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n")
	if err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		fmt.Println("Received response:", response)
	}

	response, err = client.SendMessageAndWait("{\"event\":\"getTimeStamp\",\"version\":\"test\",\"params\":{\"timestamp\":\"1702073711\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n")
	if err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		fmt.Println("Received response:", response)
	}

	//response, err = client.SendMessageAndWait("{\"event\":\"setMachines\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n")
	//if err != nil {
	//	fmt.Println("Error sending message:", err)
	//} else {
	//	fmt.Println("Received response:", response)
	//}

	//response, err = client.SendMessageAndWait("{\"event\":\"getMissions\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"date\":\"1701982800\",\"isday\":\"1\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n")
	//if err != nil {
	//	fmt.Println("Error sending message:", err)
	//} else {
	//	fmt.Println("Received response:", response)
	//}
	//
	//response, err = client.SendMessageAndWait("{\"event\":\"getMissions\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"date\":\"1701982800\",\"isday\":\"1\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n")
	//if err != nil {
	//	fmt.Println("Error sending message:", err)
	//} else {
	//	fmt.Println("Received response:", response)
	//}
	//}
	//time.Sleep(time.Second * 1)
	//common.Log(a, b)
	//
	//a, b = client.Connect.Write([]byte("{\"event\":\"setMachines\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n"))
	//a, b = client.Connect.Write([]byte("{\"event\":\"setMachines\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n"))
	//a, b = client.Connect.Write([]byte("{\"event\":\"getMissions\",\"version\":\"10.28\",\"params\":{\"user_id\":\"C2mFr8qu7ywxF5fYPxLoAQ==\",\"machine_code\":\"tI4YHjlmpnhjwhvJfPM7FA==\",\"date\":\"1701982800\",\"isday\":\"1\",\"reqid\":\"1702026206601\"},\"timestamp\":2813842600}\n"))

	//for true {
	//client.Connect.Write([]byte("{\"event\":\"getMissions\",\"params\":{\"user_id\":\"wcL+kZel294=\"},\"timestamp\":1681384045}\n"))
	//time.Sleep(time.Second * 1)
	//client.Connect.Write([]byte("{\"event\":\"getMissions\",\"version\":\"10.28\",\"params\":{\"user_id\":\"ZIYSuDf9FP/VMj5/Clw3SQ==\",\"machine_code\":\"FFoSZe0bhAh9j2FJwPtAgQ==\",\"date\":\"1698526800\"},\"timestamp\":2813842600}\n"))

	//client.Connect.Write([]byte("{\"event\":\"getNotes\",\"version\":\"10.27\",\"params\":{\"user_id\":\"wcL+kZel294=\",\"machine_code\":\"wcL+kZel294=\"},\"timestamp\":2813842600}\n"))
	//time.Sleep(time.Second * 1)

	//client.Connect.Write([]byte("{\"event\":\"getMissions\",\"version\":\"ff.11\",\"params\":{\"user_id\":\"wcL+kZel294=\",\"machine_code\":\"wcL+kZel294=\",\"date\":\"12312\"},\"timestamp\":2813842600}\n"))
	//time.Sleep(time.Second * 10)
	//client.Connect.Write([]byte("{\"event\":\"getMissions\",\"params\":{\"user_id\":\"wcL+kZel294=\",\"date\":\"12312\"},\"timestamp\":2813842600}\n"))

	////
	//client.Connect.Write([]byte("{\"Event\":\"getMissions\",\"params\":{\"user_id\":\"Rn+lqjJcpT0=\",\"date\":\"12312\"},\"timestamp\":99923333333,\"reqid\":\"99923333333\"}\n"))
	//client.Connect.Write([]byte("{\"Event\":\"getMissions\",\"params\":{\"user_id\":\"jHsnLfAZIpcoUHBZPzYWww==\",\"date\":\"12312\"},\"timestamp\":99923333333,\"reqid\":\"99923333333\"}\n"))
	//client.Connect.Write([]byte("{\"Event\":\"setMissions\",\"params\":{\"user_idf\":\"Rn+lqjJcpT0=\",\"a\":\"1233\"},\"timestamp\":1233333555}\n"))

	//}
	client.Wg.Wait()
	fmt.Println("与服务器连接断开，30秒后重试")
	time.Sleep(time.Second * 30)
	//}
}
