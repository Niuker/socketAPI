package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetMachines(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetMachines, true)
}

func GetMachines(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.GetMachines, true)
}
