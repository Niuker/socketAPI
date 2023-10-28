package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func RouterError(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.Return404, true)
}
