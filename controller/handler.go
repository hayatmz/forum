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
	
	mux.HandleFunc("/pageNewPost", pageNewPost)
	mux.HandleFunc("/formNewPost", formNewPost)
	
	mux.HandleFunc("/postLoadForm", postLoadForm)

	mux.HandleFunc("/comForm", comForm)
	mux.HandleFunc("/likeForm", likeForm)
	mux.HandleFunc("/dislikeForm", dislikeForm)
	mux.HandleFunc("/postsByLikes", postsByLikes)
	mux.HandleFunc("/postsByUser", postsByUser)

	mux.HandleFunc("/category", categoryPage)
}