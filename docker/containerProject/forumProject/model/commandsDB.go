package model

// commands to make the DB's tables
var commandsTable []string = []string{
	`CREATE TABLE IF NOT EXISTS 'users' ('id' INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, username VARCHAR(25) NOT NULL UNIQUE, password VARCHAR(40) NOT NULL, 
	email VARCHAR(320) NOT NULL UNIQUE);`,

	`CREATE TABLE IF NOT EXISTS 'posts' ('id' INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, title VARCHAR(100), 
	content VARCHAR(10000), user_id INTEGER, date datetime DEFAULT current_timestamp, FOREIGN KEY(user_id) 
	REFERENCES users(id));`,

	`CREATE TABLE IF NOT EXISTS 'categories' (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, category VARCHAR(30) UNIQUE NOT NULL);`,

	`CREATE TABLE IF NOT EXISTS 'comments' (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, user_id INTEGER, content VARCHAR(5000), post_id INTEGER, 
	date datetime DEFAULT current_timestamp, FOREIGN KEY(user_id) REFERENCES users(id), FOREIGN KEY(post_id) REFERENCES posts(id));`,

	`CREATE TABLE IF NOT EXISTS 'post_categories' (post_id INTEGER, category_id INTEGER, FOREIGN KEY(post_id) REFERENCES posts(id), FOREIGN KEY(category_id) REFERENCES categories(id), 
	PRIMARY KEY(post_id, category_id));`,

	`CREATE TABLE IF NOT EXISTS 'post_ratings' (post_id INTEGER, user_id INTEGER, rating INTEGER NOT NULL, FOREIGN KEY(post_id) REFERENCES posts(id), 
	FOREIGN KEY(user_id) REFERENCES users(id), PRIMARY KEY(post_id, user_id));`,

	`CREATE TABLE IF NOT EXISTS 'comment_ratings' (comment_id INTEGER, user_id INTEGER, rating INTEGER NOT NULL, FOREIGN KEY(comment_id) REFERENCES comments(id),
	FOREIGN KEY(user_id) REFERENCES users(id), PRIMARY KEY(comment_id, user_id));`,
}

// commands to make the DB's views
var commandsView []string = []string{viewPost, viewComment}

// command to make the  posts view
var viewPost string = `
	CREATE VIEW IF NOT EXISTS posts_view AS
		SELECT 
			p.id AS id,
			u.username AS username,
			p.title AS title,
			u.id AS user_id,
			p.content AS content,
			COALESCE(SUM(CASE WHEN pr.rating = 1 THEN 1 ELSE 0 END), 0) AS likes,
			COALESCE(SUM(CASE WHEN pr.rating = 0 THEN 1 ELSE 0 END), 0) AS dislikes,
			p.date AS date
		FROM 
			posts p
		LEFT JOIN 
			post_ratings pr ON p.id = pr.post_id
		LEFT JOIN 
			users u ON p.user_id = u.id
		GROUP BY 
			p.id;
`

// command to make the comments view
var viewComment string = `
	CREATE VIEW IF NOT EXISTS comments_view AS
		SELECT 
			c.id AS id,
			p.id AS post_id,
			u.username AS username,
			u.id AS user_id,
			c.content AS content,
			COALESCE(SUM(CASE WHEN cr.rating = 1 THEN 1 ELSE 0 END), 0) AS likes,
			COALESCE(SUM(CASE WHEN cr.rating = 0 THEN 1 ELSE 0 END), 0) AS dislikes,
			c.date AS date
		FROM 
			comments c
		LEFT JOIN 
			comment_ratings cr ON c.id = cr.comment_id
		LEFT JOIN 
			users u ON c.user_id = u.id
		LEFT JOIN 
			posts p ON c.post_id = p.id
		GROUP BY 
			c.id;
`

// var commandMakeDB string = `CREATE DATABASE IF NOT EXISTS forum`