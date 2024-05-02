package controller

import (
	"net/http"
	model "forum/model"
	view "forum/view"
)

func categoryPage(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	posts, err := model.GetPostsByCategory(category)
	if err != nil {
		return
	}

	tmpl, err := view.NewTemplate("category.html")
	if err != nil {
		return
	}
	
	tmpl.Execute(w, posts.Posts)
}
