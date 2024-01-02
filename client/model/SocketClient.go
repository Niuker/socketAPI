package model

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type SocketClient struct {
	Network         string
	Address         string
	OnMessage       func(msg string)
	Connect         net.Conn
	Wg              sync.WaitGroup
	ResponseChannel chan string
}

func (client *SocketClient) Start() error {
	conn, err := net.Dial(client.Network, client.Address)
	if err != nil {
		return err
	}
	client.Wg.Add(1)
	go HandleMessage(conn, client)
	client.Connect = conn
	client.ResponseChannel = make(chan string, 1) // Buffered channel to handle one response at a time
	return nil
}

func (client *SocketClient) SendMessageAndWait(message string) (string, error) {
	_, err := client.Connect.Write([]byte(message))
	if err != nil {
		return "", err
	}

	// Wait for response
	response := <-client.ResponseChannel

	return response, nil
}

func HandleMessage(conn net.Conn, client *SocketClient) {
	var buffer [10240]byte
	msg := ""
	defer client.Wg.Done()
	for {
		readcount, err := conn.Read(buffer[:])
		if err != nil || readcount == 0 {
			fmt.Println(err)
			break
		}
		msg += string(buffer[:readcount])
		client.ResponseChannel <- string(buffer[:readcount])
		if strings.Contains(msg, "&&&&") {
			msgs := strings.Split(msg, "&&&&")
			if len(msgs) > 1 {
				for _, m := range msgs[0 : len(msgs)-1] {
					client.OnMessage(m)
				}
				msg = msgs[len(msgs)-1]
				// Send the received message to the response channel
			}
		}
	}
}
