package model


func GetPostsByCategory(category string) ([]Post, error) {
	query := `SELECT posts.id, posts.title, posts.content, posts.user_id
                FROM posts
                INNER JOIN post_categories ON posts.id = post_categories.post_id
                INNER JOIN categories ON post_categories.category_id = categories.id
                WHERE categories.category = ? 
                ORDER BY posts.date DESC`
	rows, err := db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
        if err != nil {
			return nil, err
		}
		posts = append(posts, post)
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