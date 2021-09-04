package database

import (
	"database/sql"
	"log"
)

func (dbDir Database) InsertPerson(p entity) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err := sendInsertStatement(p, db)
	if err != nil {
		log.Fatal(err)
	}

	id, err := getLastPersonID(db)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func getLastPersonID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idPessoa from Pessoas ORDER BY idPessoa DESC LIMIT 1")
	if err != nil {
		log.Fatalf("Fail to query person id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("Fail to receive person id: %s", err)
		}
	}

	return id, nil
}
