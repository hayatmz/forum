package controller

import (
	"net/http"

	"forum/model"
	view "forum/view"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	var posts model.Posts
	if posts.GetHeadersPosts(model.QueryRoot) != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
	} else {
		view.ExecTemplate(w, "index.html", "", posts.Posts)
	}
}
