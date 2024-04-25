package controller

import (
	"net/http"
	model "forum/model"
	view "forum/view"
	
)

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
}

func registerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	var email string = r.FormValue("email")
	mail := model.NewRegister(username, password, email)
	if mail != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}