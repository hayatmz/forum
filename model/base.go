package model

import (
	"fmt"

)

func Model(u, p, e string) {
	request := "INSERT INTO posts(title, content, likes, dislikes, user_id) VALUES('titreTest', 'CONTENTtetsdsdsdsdf', '5', '225', 1);"
	fmt.Println(db)
	_, err2 := db.Exec(request)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}