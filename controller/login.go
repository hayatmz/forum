package controller

import (
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

	err := model.VerifyUser(email, password)
	if err != nil {

		if err.Error() == "aller va te register" {
			http.Redirect(w, r, "/registerPage", http.StatusFound)
			return
		}

		if err.Error() == "bad informations" {
			http.Redirect(w, r, "/loginPage", http.StatusFound)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
