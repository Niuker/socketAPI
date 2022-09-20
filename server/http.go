package server

import (
	"WebsocketDemo/router"
	"github.com/gorilla/mux"
	"net/http"
)

func HttpConnect(p string) {
	muxRouter := mux.NewRouter()

	router.RegisterRoutes(muxRouter)

	server := &http.Server{
		Addr:    p,
		Handler: muxRouter,
	}
	server.ListenAndServe()
}
