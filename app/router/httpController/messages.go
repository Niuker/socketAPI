package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func AddMessages(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddMessages, true)

}

func DelMessages(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.DelMessages, true)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetMessages, true)
}
