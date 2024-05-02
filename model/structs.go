package model

type Post struct {
	ID 			string
	Title 		string
	Content 	string
	Username 	string
	UserID		int
	Likes		int
	Dislikes 	int
	Categories	[]string
	Comments	[]Comment
}


type Comment struct {
	Content		string
	Username	string
}

type Posts struct {
	Posts []Post
}
