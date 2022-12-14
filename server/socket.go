package server

import (
	"encoding/json"
	"net"
	"socketAPI/app/encr"
	"socketAPI/app/router"
	"socketAPI/app/structure"
	"socketAPI/common"
	"socketAPI/config"
	"strconv"
	"time"
)

func Socket(listion string, c map[string]map[string]chan structure.ReqData) {
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

func readConnAndSendChan(conn net.Conn, c map[string]map[string]chan structure.ReqData, mid string) {
	for {

		raw, err := common.ReadConn(conn)
		if err != nil {
			common.Log("readConnAndSendChan close", err)
			conn.Close()
			return
		}
		for _, req := range raw {
			if req.Event == "send" {
				if revicer, ok := req.Params["revicer"]; ok {
					revicerId, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, revicer)
					if err != nil || revicerId == "" {
						common.SendConn(conn, "revicerId is not reasonable")
						conn.Close()
						return
					}
					if _, ok := req.Params["context"]; ok {
						go sendContext(req, c, revicerId)
					}
				}
			}

			prodANDcons(req, conn, c, mid)

		}
	}

}
func reviceContext(conn net.Conn, c map[string]map[string]chan structure.ReqData, mid string) error {
	defer func() {
		if err := recover(); err != nil {
			common.Log("recover error reviceContext:", err)
		}
	}()
	req := structure.ReqData{}
	for {
		if c[mid] == nil {
			c[mid] = make(map[string]chan structure.ReqData)
		}
		if c[mid]["revice"] == nil {
			c[mid]["revice"] = make(chan structure.ReqData)
		}
		if req.Params != nil {
			res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
			res.Reqid = req.Reqid
			res.Data = req.Params
			resJson, err := json.Marshal(res)
			if err != nil {
				common.SendConn(conn, err.Error())
			} else {
				err = common.SendConn(conn, string(resJson))
				if err != nil {
					c[mid]["revice"] <- req
				}
			}
		}

		req = <-c[mid]["revice"]

	}

}
func sendContext(req structure.ReqData, c map[string]map[string]chan structure.ReqData, rid string) {
	c[rid]["revice"] <- req
}

func prodANDcons(req structure.ReqData, conn net.Conn, c map[string]map[string]chan structure.ReqData, mid string) {
	if c[mid] == nil {
		c[mid] = make(map[string]chan structure.ReqData)
	}
	if c[mid][req.Event] == nil {
		c[mid][req.Event] = make(chan structure.ReqData)
	}
	if req.Timestamp < int(time.Now().Add(-6*time.Second).Unix()) {
		c[mid]["timeout"] = make(chan structure.ReqData)

	}

	go router.RegisterSocketRoutes(conn, mid, c)

	if req.Timestamp < int(time.Now().Add(-6*time.Second).Unix()) {
		c[mid]["timeout"] <- req
	} else {
		c[mid][req.Event] <- req
	}
}

func handleConnection(conn net.Conn, c map[string]map[string]chan structure.ReqData) {
	defer func() {
		if err := recover(); err != nil {
			common.Log("recover error handleConnection:", err)
		}
	}()

	mid := ""
	c[mid] = nil
	raw, err := common.ReadConn(conn)
	if err != nil {
		common.Log("handleConnection close", err)
		conn.Close()
		return
	}

	if len(raw) < 1 {
		common.SendConn(conn, "Please set id before use server")
		conn.Close()
		return
	}

	if _, ok := raw[0].Params["user_id"]; !ok {
		common.SendConn(conn, "please set id before use server")
		conn.Close()
		return
	}

	mid, err = encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, string(raw[0].Params["user_id"]))
	if err != nil || mid == "" {
		common.SendConn(conn, "id is not reasonable")
		conn.Close()
		return
	}

	id, err := strconv.Atoi(mid)
	if err != nil || mid == "" {
		common.SendConn(conn, "id is not int")
		conn.Close()
		return
	}

	common.Log("login success", id)

	resFirst := structure.ResData{Data: "success set id", Timestamp: int(time.Now().Unix())}
	resJsonFirst, err := json.Marshal(resFirst)
	if err != nil {
		common.SendConn(conn, err.Error())
	} else {
		common.SendConn(conn, string(resJsonFirst))
	}
	go reviceContext(conn, c, mid)

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
