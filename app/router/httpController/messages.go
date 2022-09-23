package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func AddMessages(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddMessages)

}

func DelMessages(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.DelMessages)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetMessages)
}
