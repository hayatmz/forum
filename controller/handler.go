package controller

import (
	model "forum/model"
	"net/http"
	"strings"
)

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/", rootPage)
	
	mux.HandleFunc("/registerPage", registerPage)
	mux.Handle("/registerForm", checkFormValues(http.HandlerFunc(registerForm)))
	mux.HandleFunc("/loginPage", loginPage)
	mux.HandleFunc("/loginForm", loginForm)
	
	mux.HandleFunc("/pageNewPost", pageNewPost)
	mux.Handle("/formNewPost", checkValidSession(checkFormValues(http.HandlerFunc(formNewPost))))
	
	mux.HandleFunc("/postLoadForm", postLoadForm)
	mux.Handle("/comForm", checkValidSession(checkFormValues(http.HandlerFunc(comForm))))

	mux.Handle("/like-comment", checkValidSession(http.HandlerFunc(likeCommentForm)))
	mux.Handle("/dislike-comment", checkValidSession(http.HandlerFunc(dislikeCommentForm)))


	mux.Handle("/likeForm", checkValidSession(http.HandlerFunc(likeForm)))
	mux.Handle("/dislikeForm", checkValidSession(http.HandlerFunc(dislikeForm)))
	mux.Handle("/postsByLikes", checkValidSession(http.HandlerFunc(postsByLikes)))
	mux.Handle("/postsByUser", checkValidSession(http.HandlerFunc(postsByUser)))

	mux.HandleFunc("/category", categoryPage)
}

func checkValidSession(next http.Handler) http.Handler {
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
		r.Form.Add("idUser", idUser)
		next.ServeHTTP(w, r)
	})
}

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


// for _, values := range form {
// 	for _, value := range values {
// 		fmt.Println(values)
// 		if len(strings.TrimSpace(value)) == 0 {
			
// 		}
// 	}