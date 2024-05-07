package model

import (
    "database/sql"
    "errors"
    "github.com/gofrs/uuid"
)

// Checks if the user is already in the database
func VerifyUserLogin(email, password string) error {
    var storedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)

    if err == sql.ErrNoRows {
        return errors.New("no account associated for this email")
    } else if err != nil {
        return err
    }

    if storedPassword != password {
        return errors.New("bad password for this account")
    }

    return nil
}

func VerifyUserRegister(email, username, password string) error {
    userID, err := uuid.NewV4()
    if err != nil {
        return err
    }
    userIDString := userID.String()

    err = chekingUserInDB(email, username, password)
    if err != nil {
        return err
    }

    queryUserRegister := "INSERT INTO users(id, email, username, password) VALUES(?,?,?,?);"
    _, err = execQuery(queryUserRegister, userIDString, email, username, password)
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
