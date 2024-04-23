package model

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	const file string = "model/forum.db"
	var errDB error
	
	_, errFile := os.Stat(file)
	if errFile != nil {
		return errFile
	}

	db, errDB = sql.Open("sqlite3", file)
	if errDB != nil {
		return errDB
	}
	return nil
}