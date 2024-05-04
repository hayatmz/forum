package model

func NewRating(idUser, idPost string, rating bool) error {
	var ratingDB bool
	querySelectRating := "SELECT rating FROM post_ratings WHERE user_id = ? AND post_id = ?"
	err := db.QueryRow(querySelectRating, idUser, idPost).Scan(&ratingDB)

	if err != nil {
		return insertRating(err, idUser, idPost, rating)
	}

	err = updateRating(idUser, idPost, ratingDB, rating)
	if err != nil {
		return err
	}
	return nil
}

func updateRating(idUser, idPost string, ratingDB, newRating bool) error {
	if newRating != ratingDB {
		queryUpdateRating := "UPDATE post_ratings SET rating = ? WHERE user_id = ? AND post_id = ?"
		_, err := execQuery(queryUpdateRating, newRating, idUser, idPost)
		if err != nil {
			return err
		}
	} else {
		queryDeleteRating := "DELETE FROM post_ratings WHERE post_id = ? AND user_id = ?"
		_, err := execQuery(queryDeleteRating, idPost, idUser)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertRating(err error, idUser, idPost string, newRating bool) error {
	if err.Error() == "sql: no rows in result set" {
		queryInsertRating := "INSERT INTO post_ratings (post_id, user_id, rating) VALUES (?, ?, ?)"
		_, err = execQuery(queryInsertRating, idPost, idUser, newRating)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
