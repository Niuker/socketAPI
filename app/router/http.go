package router

import (
	"github.com/gorilla/mux"
	"socketAPI/app/router/httpController"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()
	wsRouter.HandleFunc("/missions", httpController.GetMissions).Methods("GET")
	wsRouter.HandleFunc("/missions", httpController.SetMissions).Methods("POST")
	wsRouter.HandleFunc("/timers", httpController.ControllerGetTimers).Methods("GET")
	wsRouter.HandleFunc("/timers", httpController.ControllerSetTimers).Methods("POST")
}
