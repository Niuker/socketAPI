package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetMissions(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetMissions)
}

func GetMissions(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetMissions)
}
