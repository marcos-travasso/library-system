package database

import (
	"os"
	"testing"
)

var db = Database{
	"./temp/test_db.db",
}

func TestDatabase_CreateDatabase(t *testing.T) {
	db.CreateDatabase()

	if _, err := os.Stat(db.Dir); os.IsNotExist(err) {
		t.Errorf("Fail to create database file")
	}
}
