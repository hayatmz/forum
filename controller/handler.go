package controller

import (
	"net/http"
)

func handlers(mux *http.ServeMux) { 
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)

	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
	
	mux.HandleFunc("/postPage", postPage)
	mux.HandleFunc("/postForm", postForm)
	
	mux.HandleFunc("/postLoadPage", postLoadPage)
	mux.HandleFunc("/postLoadForm", postLoadForm)

	mux.HandleFunc("/comForm", comForm)
}