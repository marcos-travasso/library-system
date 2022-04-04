package repositories

import (
	"database/sql"
	"errors"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func (dbDir Database) InsertUser(u models.User) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	personID, err := dbDir.insertPerson(u.Person)
	if err != nil {
		return 0, err
	}
	u.Person.ID = personID

	addressID, err := dbDir.insertAddress(u.Address)
	if err != nil {
		return 0, err
	}
	u.Address.ID = addressID

	err = sendStatement(u, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Usuarios", "idUsuario")
}

func (dbDir Database) SelectUser(u models.User) (models.User, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	user := models.User{}

	rows, err := db.Query(u.SQLStatement("SELECT"))
	if err != nil {
		return user, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Print(err)
		}
	}(rows)

	for rows.Next() {
		responsible := sql.NullInt32{}
		err = rows.Scan(&user.ID, &user.CellNumber, &user.PhoneNumber, &user.CPF, &user.Email, &responsible, &user.CreationDate, &user.Person.ID, &user.Person.Name, &user.Person.Gender, &user.Person.Birthday, &user.Address.ID, &user.Address.CEP, &user.Address.City, &user.Address.Neighborhood, &user.Address.Street, &user.Address.Number, &user.Address.Complement)
		if err != nil {
			return user, err
		}
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	user.CreationDate = user.CreationDate[:10]
	return user, nil
}

func (dbDir Database) SelectUsers() ([]models.User, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	userCount, err := dbDir.countRows("Usuarios")
	if err != nil {
		return nil, err
	}

	users := make([]models.User, userCount, userCount)

	rows, err := db.Query("SELECT idUsuario, celular, telefone, cpf, email, responsavel, criacao, idPessoa, nome, genero, nascimento, idEndereco, cep, cidade, bairro, rua, numero, complemento FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco)")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		responsible := sql.NullInt32{}
		err = rows.Scan(&users[i].ID, &users[i].CellNumber, &users[i].PhoneNumber, &users[i].CPF, &users[i].Email, &responsible, &users[i].CreationDate, &users[i].Person.ID, &users[i].Person.Name, &users[i].Person.Gender, &users[i].Person.Birthday, &users[i].Address.ID, &users[i].Address.CEP, &users[i].Address.City, &users[i].Address.Neighborhood, &users[i].Address.Street, &users[i].Address.Number, &users[i].Address.Complement)
		if err != nil {
			return nil, err
		}
		users[i].CreationDate = users[i].CreationDate[:10]
	}

	return users, nil
}

func (dbDir Database) DeleteUser(u models.User) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	u, err := dbDir.SelectUser(u)
	if err != nil {
		return err
	}

	err = sendStatement(u, "DELETE", db)
	if err != nil {
		return err
	}

	err = dbDir.deleteAddress(u.Address)
	if err != nil {
		return err
	}

	err = dbDir.deletePerson(u.Person)
	if err != nil {
		return err
	}

	return nil
}

func (dbDir Database) UpdateUser(u models.User) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	userIDs, err := dbDir.SelectUser(u)
	if err != nil {
		return err
	}
	u.Person.ID = userIDs.Person.ID
	u.Address.ID = userIDs.Address.ID

	err = sendStatement(u, "UPDATE", db)
	if err != nil {
		return err
	}

	err = dbDir.updateAddress(u.Address)
	if err != nil {
		return err
	}

	err = dbDir.updatePerson(u.Person)
	return err
}
