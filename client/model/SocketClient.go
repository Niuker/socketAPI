package model

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type SocketClient struct {
	Network   string
	Address   string
	OnMessage func(msg string)
	Connect   net.Conn
	Wg        sync.WaitGroup
}

func (client *SocketClient) Start() error {
	conn, err := net.Dial(client.Network, client.Address)
	if err != nil {
		return err
	}
	client.Wg.Add(1)
	go HandleMessage(conn, client)
	client.Connect = conn
	return nil
}

func HandleMessage(conn net.Conn, client *SocketClient) {
	var buffer [1024]byte //can binger
	msg := ""
	defer client.Wg.Done()
	for true {
		readcount, err := conn.Read(buffer[:])
		if err != nil || readcount == 0 {
			fmt.Println(err)
			break
		}
		msg += string(buffer[:readcount])
		fmt.Println(msg)
		if strings.Contains(msg, "&&&&") {
			msgs := strings.Split(msg, "&&&&")
			if len(msgs) > 1 {
				for _, m := range msgs[0 : len(msgs)-1] {
					client.OnMessage(m)
				}
				msg = msgs[len(msgs)-1]
			}
		}
	}
}
