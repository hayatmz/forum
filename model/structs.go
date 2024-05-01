package model

type Post struct {
	ID 			int
	Title 		string
	Content 	string
	Username 	string
	Likes		int
	Dislikes 	int
	Categories	[]string
	Comments	[]Comment
}


type Comment struct {
	Content		string
	Username	string
}