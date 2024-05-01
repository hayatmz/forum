package model

type Post struct {
	ID      int
	Title   string
	Content string
	UserID  int
}

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
