package model

import (
	"strings"
	"database/sql"
)

// update the ratio (like, dislike) of the user in relation to a post or to a comment
//
// if already existing insert it and else update it
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

// insert the ratio (like, dislike) of the user in relation to a post or a comment in the table post_ratings or comment_rating
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

// update the ratio (like, dislike) of the user in db, or delete it if he didn't have an interaction
//
// example -> delete: he has a like, he relike, so deleting from the db its interaction because he doesn't interact anymore
//
// example -> update: he has a like, he dislike, so update its interaction in the db from a like to a dislike
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

// change the query to corresponding to the post_ratings or comment_rating 
func (rt Rating) queryRT(query string) string {
	if rt.IsComment {
		return(strings.ReplaceAll(query, "post", "comment"))
	}
	return query
}