package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"socketAPI/app/structure"
	"strings"
	"time"
)

func SocketRouter(req structure.ReqData, f func(map[string]string) (interface{}, error)) structure.ResData {
	res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
	res.Reqid = req.Reqid

	if req.Event != "send" && req.Event != "revice" {

		var versions []Version
		err := Db.Select(&versions, "select * from version where name = ?", req.Event)
		Log(req.Event)
		if err != nil || len(versions) == 0 {
			res.Code = 4
			res.Error = "select version fail"
			return res
		}

		v := false
		for _, vv := range versions {
			if vv.Version == req.Version {
				v = true
				break
			}
		}

		if !v {
			res.Code = 4
			res.Error = "version not exist"
			return res
		}
	}

	data, err := f(req.Params)

	Log("socket", req, data, err)

	if err == nil {
		res.Data = data
	} else {
		res.Code = 1
		if err.Error() == "event not exist" {
			res.Code = 4
		}
		res.Error = err.Error()
	}

	return res
}

func ReadConn(conn net.Conn) ([]structure.ReqData, error) {
	//conn.SetDeadline(time.Now().Add(30 * time.Minute))

	var reqs []structure.ReqData

	buffer := make([]byte, 204800)

	n, err := conn.Read(buffer) //No3:read
	if err != nil {
		return reqs, err
	}
	readTextsArr := strings.Split(string(buffer[:n]), "\n")
	Log(readTextsArr)
	var readTexts []string
	for _, readTextArr := range readTextsArr {
		if strings.Join(strings.Fields(readTextArr), "") != "" {
			readTexts = append(readTexts, readTextArr)
		}
	}
	for _, readText := range readTexts {
		Log(conn.RemoteAddr().String(), "receive data string:", readText)
		var req structure.ReqData

		if err := json.Unmarshal([]byte(readText), &req); err != nil {
			return reqs, err
		}
		reqs = append(reqs, req)

	}

	return reqs, nil
}

func SendConn(conn net.Conn, message string, mid string, types int) error {
	var datas = []string{message, "\n"}
	message = strings.Join(datas, "")
	_, err := conn.Write([]byte(message))
	if err != nil {
		AddUserEndRecord(mid, types)
		Log(conn.RemoteAddr().String(), " connection write error: ", err)
		e := conn.Close()
		Log(conn.RemoteAddr().String(), " connection write close error: ", e)
		return err
	}
	return nil
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
