package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"strings"
)

// load and display the login page
func loginPage(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "login.html", "", nil)
}

// get the infos of the login form
//
// connect the person with a new session if its infos are correct, else return on the same page with an error message
func loginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")
	if loginEmptyCredentials(email, password) {
		view.ExecTemplate(w, "login.html", "empty credentials !", nil)
	}

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

// check if the credentials informations are empty
func loginEmptyCredentials(email, password string) bool {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return true
	}
	return false
}