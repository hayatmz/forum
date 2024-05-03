package model

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Checks if the user is already in the database
func VerifyUserLogin(email, password string) error {
	var storedPasswordHash string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPasswordHash)

	if err == sql.ErrNoRows {
		return errors.New("no account associated with this email")
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return errors.New("bad password for this account")
	}

	return nil
}

func VerifyUserRegister(email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = checkingUserInDB(email, username)
	if err != nil {
		return err
	}

	queryUserRegister := "INSERT INTO users(email, username, password) VALUES(?,?,?);"
	_, err = execQuery(queryUserRegister, email, username, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func checkingUserInDB(email, username string) error {
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
