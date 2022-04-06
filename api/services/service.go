package services

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

var db *sql.DB
var Mock sqlmock.Sqlmock

func InitializeServices() {
	db = repositories.InitializeDatabase()
}

func InitializeTestServices() {
	var err error
	db, Mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
}
