package router

import (
	"github.com/gorilla/mux"
	"socketAPI/app/router/httpController"
	"socketAPI/common"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()

	var eventController []common.EventController
	err := common.Db.Select(&eventController, "select * from eventcontroller where disable = ?", 0)
	if err != nil {
		common.Log("get eventcontroller  error", err)
	}
	for _, ev := range eventController {
		var versions []common.Version
		err := common.Db.Select(&versions, "select * from version where name = ?", ev.Event)
		if err != nil {
			common.Log("get versions error", err)
			continue
		}
		if len(versions) == 0 {
			common.Log("get versions eq 0", err)
			continue
		}
		for _, vv := range versions {
			wsRouter.HandleFunc("/"+ev.Event, httpController.RouterError).Methods("POST")
			if vv.Version != "" {
				switch ev.Name {
				case "getMissions":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetMissionsWithMachine).Methods("POST")
				case "setMissions":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.SetMissionsWithMachine).Methods("POST")
				case "getTimers":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetTimersWithMachine).Methods("POST")
				case "setTimers":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.SetTimersWithMachine).Methods("POST")
				case "getMachines":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetMachines).Methods("POST")
				case "setMachines":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.SetMachines).Methods("POST")
				case "questions":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.UploadTQuestion).Methods("POST")
				case "addNotes":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetNotes).Methods("POST")
				case "getNotes":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.AddNotes).Methods("POST")
				case "getSystemTime":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetSystemTimers).Methods("POST")
				case "userRecord":
					wsRouter.HandleFunc("/"+ev.Event+"/"+vv.Version, httpController.GetUserRecord).Methods("POST")
				}
			}
		}
	}

	wsRouter.HandleFunc("/system/timestamp", httpController.GetSystemTimers).Methods("GET")

	wsRouter.HandleFunc("/upload", httpController.UploadPic1).Methods("POST")
	wsRouter.HandleFunc("/version", httpController.GetVersion).Methods("GET")

	wsRouter.HandleFunc("/account", httpController.Account).Methods("POST")
	wsRouter.HandleFunc("/config", httpController.GetConfig).Methods("GET")
	wsRouter.HandleFunc("/config", httpController.AddUserConfig).Methods("POST")
	wsRouter.HandleFunc("/del/config", httpController.DelConfig).Methods("POST")

	wsRouter.HandleFunc("/cron/uids", httpController.AddUids).Methods("POST")
	wsRouter.HandleFunc("/cron/uids", httpController.GetUids).Methods("GET")
	wsRouter.HandleFunc("/cron/del/uids", httpController.DelUids).Methods("POST")

	wsRouter.HandleFunc("/cron/gifts", httpController.AddGifts).Methods("POST")
	wsRouter.HandleFunc("/cron/gifts", httpController.GetGifts).Methods("GET")
	wsRouter.HandleFunc("/cron/del/gifts", httpController.DelGifts).Methods("POST")

	wsRouter.HandleFunc("/{rest:.*?}", httpController.RouterError).Methods("POST")

}
