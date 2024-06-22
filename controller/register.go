package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

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
	r.ParseForm()
	var email string = r.FormValue("email")
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		return
	}

	err = model.VerifyUserRegister(email, username, string(hashedPassword))
	if err != nil {
		if err.Error() == "email already taken" {
			http.Redirect(w, r, "/loginPage", http.StatusFound)
		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}