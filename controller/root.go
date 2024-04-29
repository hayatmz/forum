package controller

import (
	"fmt"
	"net/http"

	"forum/model"
	view "forum/view"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	var posts model.Posts
	posts.PostsRoot()
	fmt.Println(posts)
	tmpl, _ := view.NewTemplate("index.html")
	tmpl.Execute(w, posts.Posts)
}
