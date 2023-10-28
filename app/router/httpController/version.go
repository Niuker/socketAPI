package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetVersions, false)
}
