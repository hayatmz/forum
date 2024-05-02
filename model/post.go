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

// Check if the category is existing in the table categories.
//
// If the category exist return its id.
//
// If the category doesn't exist, insert it in the table category,
// and return its id.
func getIdCategory(category string) (int64, error) {
	var err error
	var idCategory int64

	queryCategory := "SELECT id FROM categories WHERE category = ?"
	db.QueryRow(queryCategory, category).Scan(&idCategory)

	if idCategory == 0 {
		queryCategory := "INSERT INTO categories (category) VALUES (?)"
		idCategory, err = execQuery(queryCategory, category)
		if err != nil {
			return 0, err
		}
	}
	return idCategory, err
}

func (post *Post) LoadPost(idPost string) error {
	queryIDPost := `SELECT id, title, content, likes, dislikes,
					username FROM posts_view WHERE id = ?`
	row := db.QueryRow(queryIDPost, idPost)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Username)
	if err != nil {
		fmt.Println(err)
	}

	queryCategories := `SELECT categories.category FROM categories 
						INNER JOIN post_categories ON categories.id = post_categories.category_id 
						WHERE post_categories.post_id = ?`
	rows, err := db.Query(queryCategories, idPost)
	for rows.Next() {
		var category string
		rows.Scan(&category)
		post.Categories = append(post.Categories, category)
	}
	
	queryComments  := 	`SELECT comments.content, users.username FROM users INNER JOIN comments 
						ON users.id = comments.user_id WHERE comments.post_id = ?`

	

	
	rows, err = db.Query(queryComments, idPost)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var content string
		var username string
		rows.Scan(&content, &username)
		post.Comments = append(post.Comments, Comment{content, username})
	}
	return nil
}

func NewComment(idUser, idPost, userComment string) error {
	queryComment := "INSERT INTO `comments` (user_id, post_id, content) VALUES (?, ?, ?)"
	_, err := execQuery(queryComment, idUser, idPost, userComment)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func Rating(idUser, idPost string, rating bool) {
	var ratingDB bool
	querySelectRating := "SELECT rating FROM post_ratings WHERE user_id = ?"
	err := db.QueryRow(querySelectRating, idUser).Scan(&ratingDB)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			queryInsertRating := "INSERT INTO post_ratings (post_id, user_id, rating) VALUES (?, ?, ?)"
			_, err := execQuery(queryInsertRating, idPost, idUser, rating)
			fmt.Println(err)
		}
	} else if rating != ratingDB {
		queryUpdateRating := "UPDATE post_ratings SET rating = ? WHERE user_id = ? AND post_id = ?"
		execQuery(queryUpdateRating, rating, idUser, idPost)
	} else {
		queryDeleteRating := "DELETE FROM post_ratings WHERE post_id = ? AND user_id = ?"
		execQuery(queryDeleteRating, idPost, idUser)
	}
}

func (posts *Posts) PostsRoot() {
	var id string
	var title string
	rows, err := db.Query(`SELECT id , title FROM posts`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err := rows.Scan(&id, &title)
		var post Post
		post.Title = title
		post.ID = id
		posts.Posts = append(posts.Posts, post)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
