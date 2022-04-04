package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func (dbDir Database) insertAddress(a models.Address) (int, error) {
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

func (dbDir Database) deleteAddress(a models.Address) error {
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

func (dbDir Database) updateAddress(a models.Address) error {
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
