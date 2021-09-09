package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) InsertAuthor(a structs.Author) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	existentID, err := dbDir.checkIfRowExists(a)
	if err != nil {
		return 0, err
	}
	if existentID != 0 {
		return existentID, nil
	}

	personID, err := dbDir.insertPerson(a.Person)
	if err != nil {
		return 0, err
	}
	a.Person.ID = personID

	err = sendStatement(a, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Autores", "idAutor")
}
