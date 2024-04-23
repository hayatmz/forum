package controller

import (
	"forum/view"
	"net/http"
	"forum/model"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/", root)
	mux.HandleFunc("/register", register)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	var email string = r.FormValue("email")
	model.Model(username, password, email)
}