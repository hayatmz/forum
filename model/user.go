package model

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Checks if the user is already in the database
func VerifyUserLogin(email, password string) (int, error) {
	var storedPasswordHash string
	var idUser int
	err := db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&idUser, &storedPasswordHash)

	if err == sql.ErrNoRows {
		return 0, errors.New("no account associated for this email")
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
	if err != nil {
		return 0, errors.New("bad password for this account")
	}

	return idUser, nil
}

// check if the user can register him with its infos, if his email and password are not already used
// 
// if valids infos, add the user to the user table in the db
func VerifyUserRegister(email, username, password string) error {
	
	err := chekingUserInDB(email, username)
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

// check if the user can register him with its infos, if his email and password are not already used
func chekingUserInDB(email, username string) error {
	var isFieldExisting bool
	
	queryEmail := "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)"
	err := db.QueryRow(queryEmail, email).Scan(&isFieldExisting)
	if err != nil {
		return err
	}
	if isFieldExisting {
		return errors.New("email already taken")
	}

	queryUserName := "SELECT EXISTS(SELECT 1 FROM users WHERE username=?)"
	err = db.QueryRow(queryUserName, username).Scan(&isFieldExisting)
	if err != nil {
		return err
	}
	if isFieldExisting {
		return errors.New("username already taken")
	}
	return nil
}

// add a new session uuid string to a user
func NewSession(token any, idUser int) error {
	queryUser := "UPDATE `users` SET session = ? WHERE id = ?"
	_, err := execQuery(queryUser, token, idUser)
	if err != nil {
		return err
	}
	return nil
}