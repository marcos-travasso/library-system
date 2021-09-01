package database

import (
	"os"
	"testing"
)

var dir = "./test_db.db"

func TestDatabase_CreateDatabase(t *testing.T) {
	CreateDatabase(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("Fail to create database file")
	}
}
