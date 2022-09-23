package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetTimers(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetTimers)
}

func GetTimers(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetTimers)
}
