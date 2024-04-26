package controller

import (
	"net/http"

	"forum/model"
	"forum/view"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
	mux.HandleFunc("/", rootPage)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("index.html")
	tmpl.Execute(w, nil)
}

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

func registerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email string = r.FormValue("email")
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")

	err := model.VerifyUserRegister(email, username, password)
	if err != nil {
		if err.Error() == "username or email already used" {
			http.Redirect(w, r, "/LoginPage", http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
}
