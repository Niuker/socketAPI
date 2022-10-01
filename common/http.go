package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socketAPI/app/structure"
	"time"
)

func POST(w http.ResponseWriter, r *http.Request, f func(map[string]string) (interface{}, error), log bool) {
	reqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error: %s\n", err)
	}

	querys := r.PostForm
	for key, query := range querys {
		reqData[key] = query[0]
	}

	data, err := f(reqData)
	if log {
		Log("http", reqData, data)
	}

	if err != nil {
		w.WriteHeader(400)
		res.Code = 1
		res.Error = err.Error()
		msg, _ := json.Marshal(res)
		w.Write(msg)
	} else {
		res.Data = data
		msg, _ := json.Marshal(res)

		w.Write(msg)
	}
}

func GET(w http.ResponseWriter, r *http.Request, f func(map[string]string) (interface{}, error), log bool) {
	reqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string), Timestamp: int(time.Now().Unix())}
	querys := r.URL.Query()
	for key, query := range querys {
		reqData[key] = query[0]
	}
	data, err := f(reqData)
	if log {
		Log("http", reqData, data)
	}

	if err != nil {
		w.WriteHeader(400)
		res.Code = 1
		res.Error = err.Error()
		msg, _ := json.Marshal(res)
		w.Write(msg)
	} else {
		res.Data = data
		msg, _ := json.Marshal(res)
		w.Write(msg)
	}
}
