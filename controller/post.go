package controller

import (
	"forum/model"
	myFuncs "forum/myFuncs"
	view "forum/view"
	"net/http"
)

func postPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("post.html")
	tmpl.Execute(w, nil)
}

func postForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var title string = r.FormValue("title")
	var content string = r.FormValue("content")
	var categories string = r.FormValue("categories")
	var validsCategories []string = myFuncs.SliceByPrefix(categories, "#")
	
	
	model.NewPost(validsCategories, title, content, 1)
	http.Redirect(w, r, "/", http.StatusFound)

}