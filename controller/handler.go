package controller

import (
	model "forum/model"
	"net/http"
	"strings"
	"net/url"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/", rootPage)
	
	mux.HandleFunc("/registerPage", registerPage)
	mux.HandleFunc("/registerForm", registerForm)
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
	
	mux.HandleFunc("/pageNewPost", pageNewPost)
	mux.Handle("/formNewPost", userConnected(http.HandlerFunc(formNewPost)))
	
	mux.Handle("/postLoadForm", userConnected(http.HandlerFunc(postLoadForm)))
	mux.Handle("/comForm", userConnected(http.HandlerFunc(comForm)))

	mux.Handle("/like-comment", userConnected(http.HandlerFunc(likeCommentForm)))
	mux.Handle("/dislike-comment", userConnected(http.HandlerFunc(dislikeCommentForm)))


	mux.Handle("/likeForm", userConnected(http.HandlerFunc(likeForm)))
	mux.Handle("/dislikeForm", userConnected(http.HandlerFunc(dislikeForm)))
	mux.Handle("/postsByLikes", userConnected(http.HandlerFunc(postsByLikes)))
	mux.Handle("/postsByUser", userConnected(http.HandlerFunc(postsByUser)))

	mux.HandleFunc("/category", categoryPage)
}

func userConnected(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusUnauthorized)
			return
		}
		idUser, err := model.GetIdUser(cookie.Value)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return
		}

		r.ParseForm()
		if !checkFormValues(r.Form) {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return
		}
		r.Form.Add("idUser", idUser)
		next.ServeHTTP(w, r)
	})
}

func checkFormValues(form url.Values) bool {
	for _, values := range form {
		for _, value := range values {
			if len(strings.TrimSpace(value)) == 0 {
				return false
			}
		}
	}
	return true
}