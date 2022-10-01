package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func UploadQuestion(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.UploadQuestion, false)
}
