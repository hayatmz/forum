package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"strconv"
)

func categoryPage(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category")
	categoryIdINT, err := strconv.Atoi(categoryID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusBadRequest)
	} else {

		var posts model.Posts
		if posts.GetHeadersPosts(model.QueryCategories, categoryIdINT) != nil {
			w.WriteHeader(http.StatusBadRequest)
			view.ExecTemplate(w, "error.html", http.StatusBadRequest)
		} else {
			view.ExecTemplate(w, "headers", posts.Posts)
		}
	}
}
