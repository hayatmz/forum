package controller

import (
	"forum/view"
	"net/http"
	"forum/model"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("index.html")
	tmpl.Execute(w, nil)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
}

func registerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	var email string = r.FormValue("email")
	model.Model(username, password, email)
	http.Redirect(w, r, "https://www.google.com/", http.StatusOK)
}