package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DatabaseDirectory = "./repositories/database.db"

func InitializeDatabase() *sql.DB {
	conn, err := sql.Open("sqlite3", DatabaseDirectory)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	return conn
}
