package controller

import (
	"net/http"
)

func comForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
}