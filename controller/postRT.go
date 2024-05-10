package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"strconv"
)

func likeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("idUser")
	idPost := r.FormValue("id-post")

	var rt model.Rating = model.Rating{idUser, idPost, false, true}
	err := rt.NewRating()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func dislikeForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("idUser")
	idPost := r.FormValue("id-post")

	var rt model.Rating = model.Rating{idUser, idPost, false, false}
	err := rt.NewRating()
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func postsByLikes(w http.ResponseWriter, r *http.Request) {
	var idUser string = r.FormValue("idUser")
	idUserINT, err := strconv.Atoi(idUser)
	if err != nil {
		view.ExecTemplate(w, "error.html", http.StatusBadRequest)
		return
	}

	var posts model.Posts
	err = posts.GetHeadersPosts(model.QueryLikes, idUserINT)
	if err != nil || posts.Posts == nil {
		w.WriteHeader(http.StatusInternalServerError)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		view.ExecTemplate(w, "headers", posts.Posts)
	}
}