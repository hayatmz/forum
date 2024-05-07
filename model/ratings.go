package model

import (
	"strings"
	"database/sql"
)

func (rt Rating) NewRating() error {
	var ratingDB bool
	querySelectRating := rt.queryRT("SELECT rating FROM post_ratings WHERE user_id = ? AND post_id = ?")
	err := db.QueryRow(querySelectRating, rt.IdUser, rt.IdPost).Scan(&ratingDB)

	if err != nil {
		return rt.insertRating(err)
	}

	err = rt.updateRating(ratingDB)
	if err != nil {
		return err
	}
	return nil
}

func (rt Rating) insertRating(err error) error {
	if err == sql.ErrNoRows {
		queryInsertRating := rt.queryRT("INSERT INTO post_ratings (post_id, user_id, rating) VALUES (?, ?, ?)")
		_, err = execQuery(queryInsertRating, rt.IdPost, rt.IdUser, rt.NewRT)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func (rt Rating) updateRating(ratingDB bool) error {
	if rt.NewRT != ratingDB {
		queryUpdateRating := rt.queryRT("UPDATE post_ratings SET rating = ? WHERE user_id = ? AND post_id = ?")
		_, err := execQuery(queryUpdateRating, rt.NewRT, rt.IdUser, rt.IdPost)
		if err != nil {
			return err
		}
	} else {
		queryDeleteRating := rt.queryRT("DELETE FROM post_ratings WHERE post_id = ? AND user_id = ?")
		_, err := execQuery(queryDeleteRating, rt.IdPost, rt.IdUser)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rt Rating) queryRT(query string) string {
	if rt.IsComment {
		return(strings.ReplaceAll(query, "post", "comment"))
	}
	return query
}