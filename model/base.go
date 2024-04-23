package model

import (
	"errors"
	"fmt"
	"log"
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
	
}