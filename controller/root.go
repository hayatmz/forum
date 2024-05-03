package controller

import (
	"net/http"

	"forum/model"
	view "forum/view"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	var posts model.Posts
	posts.GetHeadersPosts(model.QueryRoot)
	tmpl, _ := view.NewTemplate("index.html")
	tmpl.Execute(w, posts.Posts)
}
