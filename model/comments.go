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
	queryComments  := 	`SELECT comments.content, users.username FROM users INNER JOIN comments 
						ON users.id = comments.user_id WHERE comments.post_id = ?`
	rows, err := db.Query(queryComments, idPost)
	if err != nil {
		return nil
	}

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Content, &comment.Username)
		if err == nil {
			comments = append(comments, comment)
		}
	}
	return nil
}