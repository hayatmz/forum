package model

const (
	QueryCategories = `SELECT posts_view.id, posts_view.title, posts_view.username
	FROM posts_view INNER JOIN post_categories 
	ON posts_view.id = post_categories.post_id
	WHERE post_categories.category_id = ? 
	ORDER BY posts_view.date DESC`
	
	QueryRoot = `SELECT id, title, username FROM posts_view`

	QueryLikes = `SELECT posts_view.id, posts_view.title, posts_view.username FROM posts_view INNER JOIN 
	post_ratings ON posts_view.id = post_ratings.post_id WHERE post_ratings.user_id = ? AND post_ratings.rating = 1`

	QueryUserPosts = `SELECT id, title, username FROM posts_view WHERE user_id = ?`
)

// Prepare the query which returns a stmt.
//
// With the stmt, execute the request with the args which return a sql result and with this one,
// return the id of this sql Result.
func execQuery(query string, args ...any) (int64, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	
	idRes, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return idRes, nil
}