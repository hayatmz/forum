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
	var storedEmail string
	var storedUsername string
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&storedEmail)
	err2 := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&storedUsername)

	if err != nil {
		fmt.Println("err", err)
		return err
	}
	if err2 != nil {
		fmt.Println("err2", err2)
		return err
	}

	if email != storedEmail && username != storedUsername {
		_, err3 := db.Exec("INSERT INTO users(email, username, password) VALUES(?,?,?);", email, username, password)
		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
	} else {
		return fmt.Errorf("username or email already used")
	}
	return nil
}
