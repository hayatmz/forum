package controller

import (
	"fmt"
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
	
	
	err := model.NewPost(validsCategories, title, content, 1)
	if err != nil {
		http.Redirect(w, r, "/postPage", http.StatusFound)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func postLoadPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := view.NewTemplate("postLoad.html")
	tmpl.Execute(w, nil)	
}

func postLoadForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var idPost string = r.FormValue("idPost")
	LoadUniquePage(w, idPost)
}

func LoadUniquePage(w http.ResponseWriter, idPost string) {
	var post model.Post
	post.LoadPost(idPost)
	tmpl, err := view.NewTemplate("postUnique.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, post)
}