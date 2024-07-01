package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

// load and display the register's page
func registerPage(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "register.html", "", nil)
}

// Parse the form and take his informations.
//
// Hash the password with the default cost.
//
// Use the VerifyUserRegister for storing users informations in the database.
//
// Redirect to the login page if the email is already taken.
func registerForm(w http.ResponseWriter, r *http.Request) {
	email, username, hashedPassword, err := getUserInfos(r)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		return
	}

	err = model.VerifyUserRegister(email, username, string(hashedPassword))
	if err != nil {
		if err.Error() == "email already taken" {
			view.ExecTemplate(w, "login.html", err.Error(), nil)
		} else if err.Error() == "username already taken" {
			view.ExecTemplate(w, "login.html", err.Error(), nil)
		} else {
			view.ExecTemplate(w, "login.html", err.Error(), nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// get the infos from the request's form and return the same infos with the hashed password
func getUserInfos(r *http.Request) (string, string, []byte, error) {
	r.ParseForm()
	var email string = r.FormValue("email")
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return email, username, hashedPassword, err
}
