package controller

import (
	"net/http"
	model "forum/model"
	view "forum/view"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "login.html", nil)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	err := model.VerifyUserLogin(email, password)
	if err != nil {

		if err.Error() == "no account associated for this email" {
			http.Redirect(w, r, "/registerPage", http.StatusFound)
		}

		if err.Error() == "bad password for this account" {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

