package model

import (
	"fmt"
	"strconv"
)


func GetPostsByCategory(category string) (Posts, error) {
	IDCategory, _ := strconv.Atoi(category)
	query := `SELECT posts_view.id, posts_view.title, posts_view.username
                FROM posts_view
                INNER JOIN post_categories ON posts_view.id = post_categories.post_id
                WHERE post_categories.category_id = ?
                ORDER BY posts_view.date DESC`
	rows, err := db.Query(query, IDCategory)
	if err != nil {
		fmt.Println(err)
		return Posts{}, err
	}
	defer rows.Close()
	var posts Posts
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Username)
		post.Categories = getCategoriesPost(post.ID)
		fmt.Println(err)
        if err != nil {
			return Posts{}, err
		}
		posts.Posts = append(posts.Posts, post)
	}
	return posts, nil
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

func getCategoriesPost(idPost string) []Category {
	queryCategories := `SELECT categories.id, categories.category FROM categories INNER JOIN post_categories 
	ON categories.id = post_categories.category_id WHERE post_categories.post_id = ?`
	rows, err := db.Query(queryCategories, idPost)
	fmt.Println(err)
	var categories []Category
	var category Category
	for rows.Next() {
		rows.Scan(&category.IDCategory, &category.Category)
		categories = append(categories, category)
	}
	return categories
}