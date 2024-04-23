package model

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

var db *sql.DB

func Model(u, p, e string) {
	request := "INSERT INTO users(username, password, email) VALUES('homer', 'bryaxc@gmail.com', 'qdwxax');"
	fmt.Println(db)
	r, err2 := db.Exec(request)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(r.RowsAffected())

}