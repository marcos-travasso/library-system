package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

type entity interface {
	SQLStatement(statementType string) (string, error)
}

func initializeDatabase(dbDir Database) *sql.DB {
	dbDir.CreateDatabase()
	conn, err := sql.Open("sqlite3", dbDir.Dir)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	return conn
}

func (dbDir Database) InsertUser(u structs.User) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	personID, err := dbDir.InsertPerson(u.Person)
	if err != nil {
		log.Fatal(err)
	}
	u.Person.ID = personID

	addressID, err := dbDir.insertAddress(u.Address)
	if err != nil {
		log.Fatal(err)
	}
	u.Address.ID = addressID

	err = sendInsertStatement(u, db)
	if err != nil {
		log.Fatal(err)
	}

	id, err := getLastUserID(db)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (dbDir Database) SelectUser(u entity) (structs.User, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	user := structs.User{}

	rows, err := db.Query(u.SQLStatement("SELECT"))
	if err != nil {
		log.Fatalf("Fail to query user id: %s", err)
		return user, err
	}

	for rows.Next() {
		responsible := sql.NullInt32{}
		err = rows.Scan(&user.ID, &user.Person.ID, &user.CellNumber, &user.PhoneNumber, &user.Address.ID, &user.CPF, &user.Email, &responsible, &user.CreationDate, &user.Person.ID, &user.Person.Name, &user.Person.Gender, &user.Person.Birthday, &user.Address.ID, &user.Address.CEP, &user.Address.City, &user.Address.Neighborhood, &user.Address.Street, &user.Address.Number, &user.Address.Complement)
		if err != nil {
			log.Fatalf("Fail to receive user id: %s", err)
		}
	}

	user.CreationDate = user.CreationDate[:10]

	return user, nil
}

func getLastUserID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idUsuario from Usuarios ORDER BY idUsuario DESC LIMIT 1")
	if err != nil {
		log.Fatalf("Fail to query user id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("Fail to receive user id: %s", err)
		}
	}

	return id, nil
}

func getLastAddressID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idEndereco from Enderecos ORDER BY idEndereco DESC LIMIT 1")
	if err != nil {
		log.Fatalf("Fail to query address id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("Fail to receive address id: %s", err)
		}
	}

	return id, nil
}

func (dbDir Database) insertAddress(a structs.Address) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err := sendInsertStatement(a, db)
	if err != nil {
		log.Fatal(err)
	}

	id, err := getLastAddressID(db)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}
