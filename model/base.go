package model

import (
	"database/sql"
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

// Checks if the user is already in the database
func VerifyUser(email, password string) error {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)

	if err == sql.ErrNoRows {
		return fmt.Errorf("aller va te register")
	} else if err != nil {
		return err
	}

	if storedPassword != password {
		return fmt.Errorf("bad informations")
	}

	return nil
}

func VerifyUserRegister(email, username, password string) error {
	_, err3 := db.Exec("INSERT INTO users(email, username, password) VALUES(?,?,?);", email, username, password)

	return err3
}
