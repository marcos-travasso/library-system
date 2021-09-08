package database

import (
	"database/sql"
	"log"
)

func (dbDir Database) insertPerson(p entity) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(p, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Pessoas", "idPessoa")
}

func (dbDir Database) deletePerson(p entity) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(p, "DELETE", db)
	if err != nil {
		return err
	}

	return nil
}

func (dbDir Database) updatePerson(p entity) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(p, "UPDATE", db)
	if err != nil {
		return err
	}

	return nil
}
