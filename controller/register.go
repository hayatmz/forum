package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
)

func registerPage(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "register.html", "", nil)
}


func registerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email string = r.FormValue("email")
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")

	err := model.VerifyUserRegister(email, username, password)
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