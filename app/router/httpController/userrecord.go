package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func GetUserRecord(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetUserRecord, true)
}
