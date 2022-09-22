package httpController

import (
	"WebsocketDemo/app/services"
	"WebsocketDemo/app/structure"
	"encoding/json"
	"fmt"
	"net/http"
)

func ControllerSetTimers(w http.ResponseWriter, r *http.Request) {
	setTimersReqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string)}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error: %s\n", err)
	}

	querys := r.PostForm
	for key, query := range querys {
		setTimersReqData[key] = query[0]
	}

	if id, ok := querys["user_id"]; ok {
		setTimersReqData["user_id"] = id[0]
	}

	data, err := services.SetTimers(setTimersReqData)
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

func ControllerGetTimers(w http.ResponseWriter, r *http.Request) {
	getTimersReqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string)}
	query := r.URL.Query()

	if id, ok := query["user_id"]; ok {
		getTimersReqData["user_id"] = id[0]
	}

	data, err := services.GetTimers(getTimersReqData)
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
