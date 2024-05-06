package controller

import (
	"fmt"
	model "forum/model"
	view "forum/view"
	"net/http"
)

func comForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")
	userComment := r.FormValue("user-com")

	err := model.NewComment(idUser, idPost, userComment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func likeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")

	err := model.NewRating(idUser, idPost, true)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func dislikeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")

	err := model.NewRating(idUser, idPost, false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func postsByLikes(w http.ResponseWriter, r *http.Request) {

	var posts model.Posts
	err := posts.GetHeadersPosts(model.QueryLikes, 3)
	if err != nil || posts.Posts == nil {
		w.WriteHeader(http.StatusInternalServerError)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		view.ExecTemplate(w, "headers", posts.Posts)
	}
}