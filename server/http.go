package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"socketAPI/app/router"
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
