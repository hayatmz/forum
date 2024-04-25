package controller

import (
	"net/http"
	view "forum/view"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("index.html")
	tmpl.Execute(w, nil)
}