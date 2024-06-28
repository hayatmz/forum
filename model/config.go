package model

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// open the DB and return an error if the it's not possible
func InitDB() error {
	const file string = "model/database/forum.db"

	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	db, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	return(initCommandsDB())
}


// make the tables and the views of the DB
func initCommandsDB() error {
	for _, commandTable := range commandsTable {
		_, err := execQuery(commandTable)
		if err != nil {
			return err
		}
	}

	for _, commandView := range commandsView {
		_, err := execQuery(commandView)
		if err != nil {
			return err
		}
	}
	return nil
}


