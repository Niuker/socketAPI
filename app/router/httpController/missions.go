package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetMissionsWithMachine(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetMissionsWithMachine, true)
}

func GetMissionsWithMachine(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetMissionsWithMachine, true)
}
