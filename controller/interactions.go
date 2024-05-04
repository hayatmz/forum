package controller

import (
	model "forum/model"
	"net/http"
)

func comForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")
	userComment := r.FormValue("user-com")
	model.NewComment(idUser, idPost, userComment)
	LoadUniquePage(w, idPost)
}

func likeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")
	model.NewRating(idUser, idPost, true)
	LoadUniquePage(w, idPost)
}

func dislikeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("id-user")
	idPost := r.FormValue("id-post")
	model.NewRating(idUser, idPost, false)
	LoadUniquePage(w, idPost)
}