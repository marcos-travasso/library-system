package repositories

import (
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertUser(u models.User) (int64, error) {
	db := initializeDatabase()
	defer db.Close()

	//TODO maybe pass this insertperson and address to service
	personID, err := InsertPerson(u.Person)
	if err != nil {
		log.Println("insert person error: " + err.Error())
		return 0, err
	}
	u.Person.ID = int(personID)

	addressID, err := InsertAddress(u.Address)
	if err != nil {
		log.Println("insert address error: " + err.Error())
		return 0, err
	}
	u.Address.ID = int(addressID)

	result, err := db.Exec("INSERT INTO Usuarios(pessoa, celular, telefone, endereco, cpf, email, criacao) values (?, ?, ?, ?, ?, ?, date('now'))", u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email)
	if err != nil {
		log.Println("insert user error: " + err.Error())
		return 0, err
	}

	return result.LastInsertId()
}

func SelectUser(u models.User) (models.User, error) {
	db := initializeDatabase()
	defer db.Close()

	var user models.User

	//TODO check if responsible is necessary
	row := db.QueryRow("SELECT idUsuario, celular, telefone, cpf, email, criacao, idPessoa, nome, genero, nascimento, idEndereco, cep, cidade, bairro, rua, numero, complemento FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco) WHERE idUsuario =?", u.ID)
	if row.Err() != nil {
		log.Println("select user error: " + row.Err().Error())
		return user, row.Err()
	}

	err := row.Scan(&user.ID, &user.CellNumber, &user.PhoneNumber, &user.CPF, &user.Email, &user.CreationDate, &user.Person.ID, &user.Person.Name, &user.Person.Gender, &user.Person.Birthday, &user.Address.ID, &user.Address.CEP, &user.Address.City, &user.Address.Neighborhood, &user.Address.Street, &user.Address.Number, &user.Address.Complement)
	if err != nil {
		log.Println("scan user error: " + err.Error())
		return user, err
	}

	user.CreationDate = user.CreationDate[:10]
	return user, nil
}

func SelectUsers() ([]models.User, error) {
	db := initializeDatabase()
	defer db.Close()

	users := make([]models.User, 0)

	rows, err := db.Query("SELECT idUsuario, celular, telefone, cpf, email, criacao, idPessoa, nome, genero, nascimento, idEndereco, cep, cidade, bairro, rua, numero, complemento FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco)")
	if err != nil {
		log.Println("select users error: " + err.Error())
		return nil, err
	}

	for rows.Next() {
		var newUser models.User

		err = rows.Scan(&newUser.ID, &newUser.CellNumber, &newUser.PhoneNumber, &newUser.CPF, &newUser.Email, &newUser.CreationDate, &newUser.Person.ID, &newUser.Person.Name, &newUser.Person.Gender, &newUser.Person.Birthday, &newUser.Address.ID, &newUser.Address.CEP, &newUser.Address.City, &newUser.Address.Neighborhood, &newUser.Address.Street, &newUser.Address.Number, &newUser.Address.Complement)
		if err != nil {
			log.Println("scan user error: " + err.Error())
			return nil, err
		}
		newUser.CreationDate = newUser.CreationDate[:10]

		users = append(users, newUser)
	}

	return users, nil
}

func DeleteUser(u models.User) error {
	//TODO
	return nil
}

func UpdateUser(u models.User) error {
	//TODO
	return nil
}
