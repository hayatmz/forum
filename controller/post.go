package controller

import (
	"forum/model"
	myFuncs "forum/myFuncs"
	view "forum/view"
	"net/http"
)

func pageNewPost(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "post.html", nil)
}

func formNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var title string = r.FormValue("title")
	var content string = r.FormValue("content")
	var categories string = r.FormValue("categories")
	var validsCategories []string = myFuncs.SliceByPrefix(categories, "#")

	err := model.NewPost(validsCategories, title, content, 1)
	if err != nil {
		view.ExecTemplate(w, "error.html", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func postLoadForm(w http.ResponseWriter, r *http.Request) {
	var idPost string = r.URL.Query().Get("id-post")

	var post model.Post
	err := post.LoadPost(idPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", http.StatusBadRequest)
	} else {
		view.ExecTemplate(w, "postUnique.html", post)
	}
}

func postsByUser(w http.ResponseWriter, r *http.Request) {
	var posts model.Posts
	err := posts.GetHeadersPosts(model.QueryUserPosts, 2)
	if err != nil || posts.Posts == nil {
		w.WriteHeader(http.StatusInternalServerError)
		view.ExecTemplate(w, "error.html", http.StatusInternalServerError)
	} else {
		view.ExecTemplate(w, "headers", posts.Posts)
	}
}
