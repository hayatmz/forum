package model

import "fmt"

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
	queryIDPost := `SELECT posts.title, posts.content, posts.likes, posts.dislikes,
					users.username FROM posts LEFT JOIN users ON posts.user_id = users.id WHERE posts.id = ?`
	row := db.QueryRow(queryIDPost, idPost)
	err := row.Scan(&post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Username)
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
	
	// queryComments  := `SELECT content FROM comments WHERE post_id = ?`
	return nil
}

func NewComment() {

}