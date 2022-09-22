package router

import (
	"WebsocketDemo/services"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/").Subrouter()
	wsRouter.HandleFunc("/missions", services.ControllerGetMissions).Methods("GET")
	wsRouter.HandleFunc("/missions", services.ControllerSetMissions).Methods("POST")
	wsRouter.HandleFunc("/timers", services.ControllerGetTimers).Methods("GET")
	wsRouter.HandleFunc("/timers", services.ControllerSetTimers).Methods("POST")
}
