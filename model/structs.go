package model

type Post struct {
	ID 			string
	Title 		string
	Content 	string
	Username 	string
	UserID		int
	Likes		int
	Dislikes 	int
	Categories	[]Category
	Comments	[]Comment
}

type Rating struct {
	IdUser 		string
	IdPost 		string
	IsComment	bool
	NewRT 		bool
}

type Comment struct {
	ID			int
	Content		string
	Username	string
	Likes		string
	Dislikes	string
}

type Posts struct {
	Posts []Post
}

type Category struct {
	Category	string
	IDCategory	string
}