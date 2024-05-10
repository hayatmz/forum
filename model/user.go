package model

import (
	"database/sql"
	"errors"
)

// Checks if the user is already in the database
func VerifyUserLogin(email, password string) (int, error) {
	var storedPassword string
	var idUser int
	err := db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&idUser, &storedPassword)

	if err == sql.ErrNoRows {
		return 0, errors.New("no account associated for this email")
	} else if err != nil {
		return 0, err
	}

	if storedPassword != password {
		return 0, errors.New("bad password for this account")
	}

	return idUser, nil
}


func VerifyUserRegister(email, username, password string) error {
	
	err := chekingUserInDB(email, username, password)
	if err != nil {
		return err
	}

	queryUserRegister := "INSERT INTO users(email, username, password) VALUES(?,?,?);"
	_, err = execQuery(queryUserRegister, email, username, password)
	if err != nil {
		return err
	}
	return nil
}

func chekingUserInDB(email, username, password string) error {
	var isFieldExisting bool
	
	queryEmail := "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)"
	err := db.QueryRow(queryEmail, email).Scan(&isFieldExisting)
	if err != nil {
		return err
	}
	if isFieldExisting {
		return errors.New("email already taken")
	}

	queryUserName := "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)"
	err = db.QueryRow(queryUserName, password).Scan(&isFieldExisting)
	if err != nil {
		return err
	}
	if isFieldExisting {
		return errors.New("username already taken")
	}
	return nil
}

func NewSession(token string, idUser int) error {
	queryUser := "UPDATE `users` SET session = ? WHERE id = ?"
	_, err := execQuery(queryUser, token, idUser)
	if err != nil {
		return err
	}
	return nil
}