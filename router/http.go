package router

import (
	"WebsocketDemo/services"
	"WebsocketDemo/structure"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()
	wsRouter.HandleFunc("/getMissions", func(w http.ResponseWriter, r *http.Request) {
		getMissionsReqData := make(map[string]string)
		res := structure.ResData{Data: make(map[string]string)}

		getMissionsReqData["id"] = "123"
		data, err := services.GetMissions(getMissionsReqData)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		}
		res.Data = data

		msg, _ := json.Marshal(res)
		w.Write(msg)

	})

}
