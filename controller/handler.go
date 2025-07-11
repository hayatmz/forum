package controller

import (
	model "forum/model"
	"net/http"
	"strings"
)

// all handlers
func handlers(mux *http.ServeMux) {
	mux.Handle("/", isConnected(http.HandlerFunc(rootPage)))
	mux.Handle("/registerPage", isConnected(http.HandlerFunc(registerPage)))
	mux.Handle("/registerForm", checkFormValues(http.HandlerFunc(registerForm)))
	mux.Handle("/loginPage", isConnected(http.HandlerFunc(loginPage)))
	mux.Handle("/loginForm", isConnected(http.HandlerFunc(loginForm)))
	mux.Handle("/pageNewPost", isConnected(http.HandlerFunc(pageNewPost)))
	mux.Handle("/formNewPost", checkValidSession(checkFormValues(http.HandlerFunc(formNewPost))))
	mux.Handle("/postLoadForm", isConnected(http.HandlerFunc(postLoadForm)))
	mux.Handle("/comForm", checkValidSession(checkFormValues(http.HandlerFunc(comForm))))
	mux.Handle("/like-comment", checkValidSession(http.HandlerFunc(likeCommentForm)))
	mux.Handle("/dislike-comment", checkValidSession(http.HandlerFunc(dislikeCommentForm)))
	mux.Handle("/likeForm", checkValidSession(http.HandlerFunc(likeForm)))
	mux.Handle("/dislikeForm", checkValidSession(http.HandlerFunc(dislikeForm)))
	mux.Handle("/postsByLikes", checkValidSession(http.HandlerFunc(postsByLikes)))
	mux.Handle("/postsByUser", checkValidSession(http.HandlerFunc(postsByUser)))
	mux.Handle("/category", isConnected(http.HandlerFunc(categoryPage)))
	mux.Handle("/categorySearch", isConnected(http.HandlerFunc(categorySearch)))
	mux.Handle("/disconnect", checkValidSession(http.HandlerFunc(disconnect)))
}

// MIDDLEWARE
//
// check if the user has a valid session, valid session's cookie associated to a user
// 
// for actions which request to be connected as like or dislike a post...
func checkValidSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return
		}
		idUser, err := model.GetIdUser(cookie.Value)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return
		}

		r.ParseForm()
		r.Form.Add("idUser", idUser)
		w.Header().Add("connected", "connected")
		next.ServeHTTP(w, r)
	})
}

// MIDDLEWARE
//
// check if all values of a form are no empty
func checkFormValues(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for _, values := range r.Form {
			for _, value := range values {
				if len(strings.TrimSpace(value)) == 0 {
					http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

// MIDDLEWARE
//
// add to the response a header "connected" if the user is connected, valid session's cookie associated to a user
//
// this additionnal header will permit to load either the header for the connected user or for the not connected user 
func isConnected(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			goto notConnected
		}
		_, err = model.GetIdUser(cookie.Value)
		if err != nil {
			goto notConnected
		}
		r.ParseForm()
		w.Header().Add("connected", "connected")
		notConnected:
		next.ServeHTTP(w, r)
	})
}