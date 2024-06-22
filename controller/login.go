package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "login.html", "", nil)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	id, err := model.VerifyUserLogin(email, password)

	if err != nil {
		if err.Error() == "no account associated for this email" {
			view.ExecTemplate(w, "register.html", err.Error(), nil)
		} else if err.Error() == "bad password for this account" {
			view.ExecTemplate(w, "login.html", err.Error(), nil)
		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
	} else {
		newSession(w, r, id)
	}
}