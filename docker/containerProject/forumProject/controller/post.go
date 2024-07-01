package controller

import (
	"forum/model"
	myFuncs "forum/myFuncs"
	view "forum/view"
	"net/http"
	"strconv"
)

// load and display the page to post
func pageNewPost(w http.ResponseWriter, r *http.Request) {
	view.ExecTemplate(w, "post.html", "", nil)
}

// get the infos of the form to post and if no info of the form is empty, add the post of the post's list
func formNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var title string = r.FormValue("title")
	var content string = r.FormValue("content")
	var categories string = r.FormValue("categories")
	var idUser string = r.FormValue("idUser")
	
	idUserINT, err := strconv.Atoi(idUser)
	if err != nil {
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
		return
	}
	var validsCategories []string = myFuncs.SliceByPrefix(categories, "#")

	err = model.NewPost(validsCategories, title, content, idUserINT)
	if err != nil {
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// load and display the post with its infos (title, content, username, comments with likes and dislikes, likes and dislikes of the post, categories)
// load the zone to comment the post and like and dislike the comments or the post
func postLoadForm(w http.ResponseWriter, r *http.Request) {
	var idPost string = r.URL.Query().Get("id-post")

	var post model.Post
	err := post.LoadPost(idPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
	} else {
		view.ExecTemplate(w, "postUnique.html", "", post)
	}
}

// load and display the headers posts made by the connected user
func postsByUser(w http.ResponseWriter, r *http.Request) {
	var idUser string = r.FormValue("idUser")
	idUserINT, err := strconv.Atoi(idUser)
	if err != nil {
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
		return
	}
	
	var posts model.Posts
	err = posts.GetHeadersPosts(model.QueryUserPosts, idUserINT)
	if err != nil  {
		w.WriteHeader(http.StatusInternalServerError)
		view.ExecTemplate(w, "error.html", "", http.StatusInternalServerError)
	} else {
		view.ExecTemplate(w, "headers.html", "Your Posts", posts.Posts)
	}
}
