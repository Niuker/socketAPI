package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func GetUids(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetUids, true)
}

func AddUids(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddUids, true)
}

func DelUids(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.DelUids, true)
}

func GetGifts(w http.ResponseWriter, r *http.Request) {
	common.GET(w, r, services.GetGifts, true)
}

func AddGifts(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddGifts, true)
}

func DelGifts(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.DelGifts, true)
}
