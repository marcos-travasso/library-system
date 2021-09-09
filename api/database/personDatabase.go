package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) insertPerson(p structs.Person) (int, error) {
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

func (dbDir Database) deletePerson(p structs.Person) error {
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

func (dbDir Database) updatePerson(p structs.Person) error {
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
