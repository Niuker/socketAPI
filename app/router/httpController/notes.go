package httpController

import (
	"net/http"
	"socketAPI/app/services"
	"socketAPI/common"
)

func AddNotes(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.AddNotes, true)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	common.POST(w, r, services.GetNotes, true)
}
