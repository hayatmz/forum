package controller

import (
	"net/http"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)
	mux.HandleFunc("/postPage", postPage)
	mux.HandleFunc("/postForm", postForm)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
}