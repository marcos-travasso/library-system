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

	id, err := getLastPersonID(db)
	if err != nil {
		return 0, err
	}

	return id, nil
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

func getLastPersonID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idPessoa from Pessoas ORDER BY idPessoa DESC LIMIT 1")
	if err != nil {
		log.Printf("Fail to query person id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("Fail to receive person id: %s", err)
			return 0, err
		}
	}

	return id, nil
}
