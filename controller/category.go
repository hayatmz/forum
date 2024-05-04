package controller

import (
	"fmt"
	model "forum/model"
	view "forum/view"
	"net/http"
	"strconv"
)

func categoryPage(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	categoryId, _ := strconv.Atoi(category)
	var posts model.Posts
	err := posts.GetHeadersPosts(model.QueryCategories, categoryId)
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := view.NewTemplate("category.html")
	if err != nil {
		return
	}
	
	tmpl.Execute(w, posts.Posts)
}
