package model

func NewComment(idUser, idPost, userComment string) error {
	queryComment := "INSERT INTO `comments` (user_id, post_id, content) VALUES (?, ?, ?)"
	_, err := execQuery(queryComment, idUser, idPost, userComment)
	if err != nil {
		return err
	}
	return nil
}

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