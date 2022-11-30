package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func Account(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.Account, true)
}

func AddUserConfig(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddUserConfig, true)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetConfig, true)
}

func DelConfig(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.DelConfig, true)
}
