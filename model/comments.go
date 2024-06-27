package model

// add in the DB a comment in the comment table (user id, post_id, and comment content)
func NewComment(idUser, idPost, userComment string) error {
	queryComment := "INSERT INTO `comments` (user_id, post_id, content) VALUES (?, ?, ?)"
	_, err := execQuery(queryComment, idUser, idPost, userComment)
	if err != nil {
		return err
	}
	return nil
}

// get all comments of a post from the comments_view with the post id
func getCommentsPost(idPost string) []Comment {
	queryComments  := 	`SELECT id, username, content, likes, dislikes FROM comments_view WHERE comments_view.post_id = ?`
	rows, err := db.Query(queryComments, idPost)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.Username, &comment.Content, &comment.Likes, &comment.Dislikes)
		if err == nil {
			comments = append(comments, comment)
		}
	}
	return comments
}