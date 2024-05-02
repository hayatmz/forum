package model

import (
	"fmt"
)

// Insert the post with its title, content and associed isUser in the table posts.
//
// For each category existing, get its id and,
// if the category doesn't exist, insert it in the table categories and get its id.
//
// For each category, insert in the table post_categories the post id and the category id.
func NewPost(categories []string, title, content string, idUser int) error {
	queryPost := "INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)"
	idPost, err := execQuery(queryPost, title, content, idUser)
	if err != nil {
		return err
	}

	for _, category := range categories {
		idCategory, err := getIdCategory(category)
		if err != nil {
			return err
		}

		queryPostCategories := "INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)"
		_, err = execQuery(queryPostCategories, idPost, idCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func (posts *Posts) PostsRoot() {
	var id string
	var title string
	var username string
	rows, err := db.Query(`SELECT id, title, username FROM posts_view`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err := rows.Scan(&id, &title, &username)
		var post Post
		post.Categories = getCategoriesPost(id)
		post.Title = title
		post.ID = id
		post.Username = username
		posts.Posts = append(posts.Posts, post)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
