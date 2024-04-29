package model

type Post struct {
	Title 		string
	Content 	string
	Username 	string
	Likes		int
	Dislikes 	int
	Categories	[]string
}