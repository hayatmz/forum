package controller

import (
	model "forum/model"
	view "forum/view"
	"net/http"
	"strconv"
)

// filter the posts by category
func categoryPage(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category")
	categoryIdINT, err := strconv.Atoi(categoryID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
	} else {
		getHeadersCategories(w, categoryIdINT)
	}
}

// filter the posts by category from the search bar
func categorySearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idCategory, err := model.GetIdCategory(r.FormValue("categoryFilter"), false)
	if err != nil || idCategory == 0 {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		getHeadersCategories(w, int(idCategory))
	}
}

// display the headers of the posts (title, username, categories, date) 
func getHeadersCategories(w http.ResponseWriter, categoryIdINT int) {
	var posts model.Posts
		err := posts.GetHeadersPosts(model.QueryCategories, categoryIdINT)
		if err != nil || posts.Posts == nil {
			w.WriteHeader(http.StatusBadRequest)
			view.ExecTemplate(w, "error.html", "", http.StatusBadRequest)
		} else {
			view.ExecTemplate(w, "headers.html", "Posts By Category", posts.Posts)
		}
}