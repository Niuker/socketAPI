package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func UploadTQuestion(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.UploadTQuestion, false)
}
