package model

// POST INFOS
type Post struct {
	ID 			string
	Title 		string
	Content 	string
	Username 	string
	UserID		int
	Likes		int
	Dislikes 	int
	Date 		string
	Categories	[]Category
	Comments	[]Comment
}

// RATING INFOS (LIKE OR DISLIKE)
type Rating struct {
	IdUser 		string
	IdPost 		string
	IsComment	bool
	NewRT 		bool
}

// COMMENT INFOS
type Comment struct {
	ID			int
	Content		string
	Username	string
	Likes		string
	Dislikes	string
}

// REGROUPING MANY POSTS
type Posts struct {
	Posts []Post
}

// CATEGORY INFOS
type Category struct {
	Category	string
	IDCategory	string
}