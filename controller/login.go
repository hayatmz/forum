package controller

import (
	"fmt"
	"net/http"

	model "forum/model"
	view "forum/view"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("login.html")
	tmpl.Execute(w, nil)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	err := model.VerifyUserLogin(email, password)
	if err != nil {

		if err.Error() == "no account associated for this email" {
			fmt.Println("email error")
			http.Redirect(w, r, "/registerPage", http.StatusFound)
			return
		}

		if err.Error() == "bad password for this account" {
			fmt.Println("password error")
			http.Redirect(w, r, "/loginPage", http.StatusFound)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
