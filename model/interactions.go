package model

import (
	"fmt"
)

func (post *Post) LoadPost(idPost string) error {
	queryIDPost := `SELECT id, title, content, likes, dislikes,
					username FROM posts_view WHERE id = ?`
	row := db.QueryRow(queryIDPost, idPost)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Username)
	if err != nil {
		fmt.Println(err)
	}

	queryCategories := `SELECT categories.id, categories.category FROM categories 
						INNER JOIN post_categories ON categories.id = post_categories.category_id 
						WHERE post_categories.post_id = ?`
	rows, err := db.Query(queryCategories, idPost)
	for rows.Next() {
		var category Category
		rows.Scan(&category.IDCategory, &category.Category)
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
	querySelectRating := "SELECT rating FROM post_ratings WHERE user_id = ? AND post_id = ?"
	err := db.QueryRow(querySelectRating, idUser, idPost).Scan(&ratingDB)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			queryInsertRating := "INSERT INTO post_ratings (post_id, user_id, rating) VALUES (?, ?, ?)"
			execQuery(queryInsertRating, idPost, idUser, rating)
		}
	} else if rating != ratingDB {
		queryUpdateRating := "UPDATE post_ratings SET rating = ? WHERE user_id = ? AND post_id = ?"
		execQuery(queryUpdateRating, rating, idUser, idPost)
	} else {
		queryDeleteRating := "DELETE FROM post_ratings WHERE post_id = ? AND user_id = ?"
		execQuery(queryDeleteRating, idPost, idUser)
	}
}