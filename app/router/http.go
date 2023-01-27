package router

import (
	"github.com/gorilla/mux"
	"socketAPI/app/router/httpController"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()
	wsRouter.HandleFunc("/missions", httpController.GetMissionsWithMachine).Methods("GET")
	wsRouter.HandleFunc("/missions", httpController.SetMissionsWithMachine).Methods("POST")

	wsRouter.HandleFunc("/timers", httpController.GetTimersWithMachine).Methods("GET")
	wsRouter.HandleFunc("/timers", httpController.SetTimersWithMachine).Methods("POST")

	wsRouter.HandleFunc("/messages", httpController.GetMessages).Methods("GET")
	wsRouter.HandleFunc("/messages", httpController.AddMessages).Methods("POST")
	wsRouter.HandleFunc("/del/messages", httpController.DelMessages).Methods("POST")

	wsRouter.HandleFunc("/machines", httpController.GetMachines).Methods("GET")
	wsRouter.HandleFunc("/machines", httpController.SetMachines).Methods("POST")

	wsRouter.HandleFunc("/account", httpController.Account).Methods("POST")
	wsRouter.HandleFunc("/config", httpController.GetConfig).Methods("GET")
	wsRouter.HandleFunc("/config", httpController.AddUserConfig).Methods("POST")
	wsRouter.HandleFunc("/del/config", httpController.DelConfig).Methods("POST")

	wsRouter.HandleFunc("/questions", httpController.UploadQuestion).Methods("POST")
	wsRouter.HandleFunc("/t_questions", httpController.UploadTQuestion).Methods("POST")

	wsRouter.HandleFunc("/upload1", httpController.UploadPic1).Methods("POST")
	wsRouter.HandleFunc("/upload2", httpController.UploadPic2).Methods("POST")

	wsRouter.HandleFunc("/cron/uids", httpController.AddUids).Methods("POST")
	wsRouter.HandleFunc("/cron/uids", httpController.GetUids).Methods("GET")
	wsRouter.HandleFunc("/cron/del/uids", httpController.DelUids).Methods("POST")

	wsRouter.HandleFunc("/cron/gifts", httpController.AddGifts).Methods("POST")
	wsRouter.HandleFunc("/cron/gifts", httpController.GetGifts).Methods("GET")
	wsRouter.HandleFunc("/cron/del/gifts", httpController.DelGifts).Methods("POST")

}
