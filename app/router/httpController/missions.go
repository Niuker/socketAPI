package httpController

import (
	"WebsocketDemo/app/services"
	"WebsocketDemo/app/structure"
	"encoding/json"
	"fmt"
	"net/http"
)

func SetMissions(w http.ResponseWriter, r *http.Request) {
	setMissionsReqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string)}
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error: %s\n", err)
	}

	querys := r.PostForm
	for key, query := range querys {
		setMissionsReqData[key] = query[0]
	}

	if id, ok := querys["user_id"]; ok {
		setMissionsReqData["user_id"] = id[0]
	}

	if isday, ok := querys["isday"]; ok {
		setMissionsReqData["isday"] = isday[0]
	}

	if date, ok := querys["date"]; ok {
		setMissionsReqData["date"] = date[0]
	}

	data, err := services.SetMissions(setMissionsReqData)
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

func GetMissions(w http.ResponseWriter, r *http.Request) {
	getMissionsReqData := make(map[string]string)
	res := structure.ResData{Data: make(map[string]string)}
	query := r.URL.Query()

	if id, ok := query["user_id"]; ok {
		getMissionsReqData["user_id"] = id[0]
	}

	if isday, ok := query["isday"]; ok {
		getMissionsReqData["isday"] = isday[0]
	}

	if date, ok := query["date"]; ok {
		getMissionsReqData["date"] = date[0]
	}

	data, err := services.GetMissions(getMissionsReqData)
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
