package database

import (
	"os"
	"testing"
)

var db = Database{
	"./test_db.db",
}

func TestDatabase_CreateDatabase(t *testing.T) {
	db.CreateDatabase()

	if _, err := os.Stat(db.dir); os.IsNotExist(err) {
		t.Errorf("Fail to create database file")
	}
}
