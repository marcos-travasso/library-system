package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) insertGenre(g structs.Genre) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	existentID, err := dbDir.checkIfRowExists(g)
	if err != nil {
		return 0, err
	}
	if existentID != 0 {
		return existentID, nil
	}

	err = sendStatement(g, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Generos", "idGenero")
}
