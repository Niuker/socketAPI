package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func UploadPic1(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.UploadPic1, false)
}
func UploadPic2(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.UploadPic2, false)
}
