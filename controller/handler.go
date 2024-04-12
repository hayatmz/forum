package controller

import (
	"forum/view"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := view.NewTemplate()
	tmpl.Execute(w, nil)
}