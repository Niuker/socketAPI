package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func SetTimersWithMachine(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.SetTimersWithMachine, true)
}

func GetTimersWithMachine(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.GetTimersWithMachineStrict, true)
}
