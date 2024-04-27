package controller

import (
	"net/http"
	model "forum/model"
	view "forum/view"
	
)

// func registerForm(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	var username string = r.FormValue("username")
// 	var password string = r.FormValue("password")
// 	var email string = r.FormValue("email")
// 	mail := model.NewRegister(username, password, email)
// 	if mail != nil {
// 		http.Redirect(w, r, "/", http.StatusFound)
// 	}
// 	http.Redirect(w, r, "/", http.StatusFound)
// }


func registerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email string = r.FormValue("email")
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")

	err := model.VerifyUserRegister(email, username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			http.Redirect(w, r, "/registerPage", http.StatusFound)
			return
		} else if err.Error() == "UNIQUE constraint failed: users.email" {
			http.Redirect(w, r, "/loginPage", http.StatusFound)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("register.html")
	tmpl.Execute(w, nil)
}
