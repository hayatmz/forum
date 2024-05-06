package controller

import (
	"net/http"

	model "forum/model"
	view "forum/view"

	"golang.org/x/crypto/bcrypt"
)

// Execute the new template register.
func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
