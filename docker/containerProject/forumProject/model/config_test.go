package model

import (
	"testing"
)

// test init db tables and views and path file db
func TestInitDB(t *testing.T) {
	var  validPath string = "./database/forum.db"
	var invalidFile string = "./database/badDB.db"
	var invalidDirectory string = "./model/forum.db"

	err := InitDB(validPath)
	if err != nil {
		t.Fail()
	}

	err = InitDB(invalidFile)
	if err == nil {
		t.Fail()
	}

	err = InitDB(invalidDirectory)
	if err == nil {
		t.Fail()
	}
}