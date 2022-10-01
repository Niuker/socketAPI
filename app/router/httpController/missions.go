package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetMissions(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetMissions, true)
}

func GetMissions(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetMissions, true)
}
