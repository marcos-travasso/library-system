package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) insertAddress(a structs.Address) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(a, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Enderecos", "idEndereco")
}

func (dbDir Database) deleteAddress(a structs.Address) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(a, "DELETE", db)
	if err != nil {
		return err
	}

	return nil
}

func (dbDir Database) updateAddress(a structs.Address) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendStatement(a, "UPDATE", db)
	if err != nil {
		return err
	}

	return nil
}
