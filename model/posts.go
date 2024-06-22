package model

import (
	"time"
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

func (post *Post) LoadPost(idPost string) error {
	queryIDPost := `SELECT id, title, content, likes, dislikes,
					username FROM posts_view WHERE id = ?`
	row := db.QueryRow(queryIDPost, idPost)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Username)
	if err != nil {
		return err
	}
	
	post.Categories = getCategoriesPost(idPost)
	post.Comments = getCommentsPost(idPost)
	return nil
}

func (posts *Posts) GetHeadersPosts(query string, args ...any) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		var datePost time.Time
		err := rows.Scan(&post.ID, &post.Title, &post.Username, &datePost)
		convertDate(&datePost, &post.Date)
        if err != nil {
			continue
		}
		post.Categories = getCategoriesPost(post.ID)
		posts.Posts = append(posts.Posts, post)
	}
	return nil
}

func convertDate(date *time.Time, postDate *string) {
	dateFormat := date.Format("2006-01-02 15:04:05")
	*postDate = dateFormat
}