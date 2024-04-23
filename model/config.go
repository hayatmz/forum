package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	const file string = "model/forum.db"
	var errDB error
	
	db, errDB = sql.Open("sqlite", file)
	if errDB != nil {
		return nil, errDB
	}
	return db, nil
}