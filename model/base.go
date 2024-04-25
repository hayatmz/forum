package model

import (
	"fmt"
	"log"
	"database/sql"
)

type User struct {
	Id 			int
	Username 	string
	Email 		string
	Password 	string
}

func NewRegister(u, p, e string) error {
	var user User
	req := fmt.Sprintf("SELECT * FROM users WHERE email='%s'; &&", e)
	fmt.Println(e)
	rows, err := db.Query(req)
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return err
		}
		fmt.Println(user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	if user.Email == "" {
		return nil
	}
	return nil
	
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


func NewPost(categories []string, title, content string, idUser int) error {
	result, _ := db.Exec("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)", title, content, idUser)
	postId, _ := result.LastInsertId()
	fmt.Println(postId)

	var idCategory int
	for _, category := range categories {
		db.QueryRow("SELECT id FROM categories WHERE category = ?", category).Scan(&idCategory)
		if idCategory != 0 {
			_, err := db.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postId, idUser)
			fmt.Println(err)
		}
	}
	return nil
}