package server

import (
	"WebsocketDemo/common"
	"WebsocketDemo/config"
	"WebsocketDemo/encr"
	"WebsocketDemo/router"
	"WebsocketDemo/structure"
	"encoding/json"
	"net"
)

func Socket(listion string, c map[string]map[string]chan map[string]string) {
	//make socket，listen port  No1 bingding
	netListen, err := net.Listen("tcp", listion)

	common.CheckError(err)
	//defer
	defer netListen.Close()

	common.Log("Waiting for clients")
	for {
		conn, err := netListen.Accept() //NO2 get accept
		if err != nil {
			continue //err?? gun
		}

		common.Log(conn.RemoteAddr().String(), " tcp connect success")

		go handleConnection(conn, c) //goroutine gogo
	}
}

func readConnAndSendChan(conn net.Conn, c map[string]map[string]chan map[string]string, mid string) {
	for {

		raw, err := common.ReadConn(conn)
		for _, req := range raw {

			if err != nil {
				common.Log(err)
				conn.Close()
				return
			}
			if c[mid] == nil {
				c[mid] = make(map[string]chan map[string]string)
			}
			if c[mid][req.Event] == nil {
				c[mid][req.Event] = make(chan map[string]string)
			}

			go router.RegisterSocketRoutes(conn, mid, c)

			c[mid][req.Event] <- req.Params
		}
	}

}

func handleConnection(conn net.Conn, c map[string]map[string]chan map[string]string) {
	mid := ""
	c[mid] = nil
	raw, err := common.ReadConn(conn)
	if err != nil {
		common.Log(err)
		conn.Close()
		return
	}

	if _, ok := raw[0].Params["id"]; !ok {
		common.SendConn(conn, "please set id before use server")
	}

	mid, err = encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, string(raw[0].Params["id"]))
	if err != nil || mid == "" {
		common.SendConn(conn, "id is not reasonable")
	}

	resFirst := structure.ResData{Data: "success set id"}
	resJsonFirst, err := json.Marshal(resFirst)
	if err != nil {
		common.SendConn(conn, err.Error())
	} else {
		common.SendConn(conn, string(resJsonFirst))
	}

	readConnAndSendChan(conn, c, mid)

	//for {
	//	go readConn(buffer, conn)
	//
	//	if uid == "" {
	//		continue
	//	}
	//	for {
	//		words := strconv.Itoa(int(time.Now().Unix()))
	//		var datas = []string{words, "\n"}
	//		words = strings.Join(datas, "")
	//		time.Sleep(time.Second * 1)
	//		_, err := conn.Write([]byte(words))
	//		if err != nil {
	//			Log(conn.RemoteAddr().String(), " connection2 error: ", err)
	//			return
	//		}
	//	}
	//
	//	c[uid] = structure.NewChanMgr()
	//	flag := <-c[uid].C
	//
	//	if flag > 0 {
	//		words := strconv.Itoa(flag)
	//		var datas = []string{words, "\n"}
	//		words = strings.Join(datas, "")
	//
	//		_, err := conn.Write([]byte(words))
	//		Log(conn.RemoteAddr().String(), " conn write success: ", words)
	//		flag = 0
	//		if err != nil {
	//			Log(conn.RemoteAddr().String(), " connection error: ", err)
	//			return
	//		}
	//	}
	//}
}
