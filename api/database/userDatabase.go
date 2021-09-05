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
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	personID, err := dbDir.InsertPerson(u.Person)
	if err != nil {
		return 0, err
	}
	u.Person.ID = personID

	addressID, err := dbDir.insertAddress(u.Address)
	if err != nil {
		return 0, err
	}
	u.Address.ID = addressID

	err = sendInsertStatement(u, db)
	if err != nil {
		return 0, err
	}

	id, err := getLastUserID(db)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dbDir Database) SelectUser(u entity) (structs.User, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	user := structs.User{}

	rows, err := db.Query(u.SQLStatement("SELECT"))
	if err != nil {
		log.Printf("Fail to query user id: %s", err)
		return user, err
	}

	for rows.Next() {
		responsible := sql.NullInt32{}
		err = rows.Scan(&user.ID, &user.Person.ID, &user.CellNumber, &user.PhoneNumber, &user.Address.ID, &user.CPF, &user.Email, &responsible, &user.CreationDate, &user.Person.ID, &user.Person.Name, &user.Person.Gender, &user.Person.Birthday, &user.Address.ID, &user.Address.CEP, &user.Address.City, &user.Address.Neighborhood, &user.Address.Street, &user.Address.Number, &user.Address.Complement)
		if err != nil {
			log.Printf("Fail to receive user id: %s", err)
			return user, err
		}
	}

	user.CreationDate = user.CreationDate[:10]
	return user, nil
}

func (dbDir Database) SelectUsers() ([]structs.User, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	rows, err := db.Query("SELECT COUNT(*) FROM Usuarios")
	if err != nil {
		log.Printf("Fail to query user count: %s", err)
		return nil, err
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	userCount := 0
	for rows.Next() {
		err = rows.Scan(&userCount)
		if err != nil {
			log.Printf("Fail to receive user count: %s", err)
			return nil, err
		}
	}
	users := make([]structs.User, 0, userCount)

	rows, err = db.Query("SELECT * FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco)")
	if err != nil {
		log.Printf("Fail to query users: %s", err)
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		responsible := sql.NullInt32{}
		err = rows.Scan(&users[i].ID, &users[i].Person.ID, &users[i].CellNumber, &users[i].PhoneNumber, &users[i].Address.ID, &users[i].CPF, &users[i].Email, &responsible, &users[i].CreationDate, &users[i].Person.ID, &users[i].Person.Name, &users[i].Person.Gender, &users[i].Person.Birthday, &users[i].Address.ID, &users[i].Address.CEP, &users[i].Address.City, &users[i].Address.Neighborhood, &users[i].Address.Street, &users[i].Address.Number, &users[i].Address.Complement)
		if err != nil {
			log.Printf("Fail to receive users id: %s", err)
			return nil, err
		}
		users[i].CreationDate = users[i].CreationDate[:10]
	}

	return users, nil
}

func getLastUserID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idUsuario from Usuarios ORDER BY idUsuario DESC LIMIT 1")
	if err != nil {
		log.Printf("Fail to query user id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("Fail to receive user id: %s", err)
			return 0, err
		}
	}

	return id, nil
}

func getLastAddressID(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT idEndereco from Enderecos ORDER BY idEndereco DESC LIMIT 1")
	if err != nil {
		log.Printf("Fail to query address id: %s", err)
		return 0, err
	}

	id := 0

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("Fail to receive address id: %s", err)
			return 0, err
		}
	}

	return id, nil
}

func (dbDir Database) insertAddress(a structs.Address) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendInsertStatement(a, db)
	if err != nil {
		return 0, err
	}

	id, err := getLastAddressID(db)
	if err != nil {
		return 0, err
	}

	return id, nil
}
