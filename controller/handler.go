package controller

import (
	"forum/view"
	"net/http"
	"forum/model"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
	mux.HandleFunc("/", rootPage)
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
	http.Redirect(w, r, "/", http.StatusOK)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("login.html")
	tmpl.Execute(w, nil)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email string = r.FormValue("email")
	// var password string = r.FormValue("password")
	model.CheckIfInDB(email)
	http.Redirect(w, r, "/", http.StatusOK)
}