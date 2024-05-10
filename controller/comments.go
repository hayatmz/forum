package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
)

func comForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	idUser := r.FormValue("idUser")
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

func likeCommentForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("idUser")
	idComment := r.FormValue("id-comment")

	var rt model.Rating = model.Rating{idUser, idComment, true, true}
	err := rt.NewRating()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}

func dislikeCommentForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idUser := r.FormValue("idUser")
	idComment := r.FormValue("id-comment")

	var rt model.Rating = model.Rating{idUser, idComment, true, false}
	err := rt.NewRating()
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}