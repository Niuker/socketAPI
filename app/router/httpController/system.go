package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func GetSystemTimers(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.GetSystemTimers, false)
}
