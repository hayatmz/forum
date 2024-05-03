package model

func (posts *Posts) GetHeadersPosts(query string, args ...any) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		post.Categories = getCategoriesPost(post.ID)
		err := rows.Scan(&post.ID, &post.Title, &post.Username)
        if err == nil {
			posts.Posts = append(posts.Posts, post)
		}
	}
	return nil
}

const (
	QueryCategories = `SELECT posts_view.id, posts_view.title, posts_view.username
	FROM posts_view INNER JOIN post_categories 
	ON posts_view.id = post_categories.post_id
	WHERE post_categories.category_id = ? 
	ORDER BY posts_view.date DESC`
	
	QueryRoot = `SELECT id, title, username FROM posts_view`
)