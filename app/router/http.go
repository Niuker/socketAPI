package router

import (
	"github.com/gorilla/mux"
	"socketAPI/app/router/httpController"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()
	wsRouter.HandleFunc("/missions", httpController.GetMissions).Methods("GET")
	wsRouter.HandleFunc("/missions", httpController.SetMissions).Methods("POST")

	wsRouter.HandleFunc("/timers", httpController.GetTimers).Methods("GET")
	wsRouter.HandleFunc("/timers", httpController.SetTimers).Methods("POST")

	wsRouter.HandleFunc("/messages", httpController.GetMessages).Methods("GET")
	wsRouter.HandleFunc("/messages", httpController.AddMessages).Methods("POST")
	wsRouter.HandleFunc("/del/messages", httpController.DelMessages).Methods("POST")

	wsRouter.HandleFunc("/machine", httpController.GetMachines).Methods("GET")
	wsRouter.HandleFunc("/machine", httpController.SetMachines).Methods("POST")

}
