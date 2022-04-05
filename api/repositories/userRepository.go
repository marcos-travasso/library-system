package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertUser(db *sql.DB, u *models.User) (err error) {
	result, err := db.Exec("INSERT INTO Usuarios(pessoa, celular, telefone, endereco, cpf, email, criacao) values (?, ?, ?, ?, ?, ?, date('now'))", u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email)
	if err != nil {
		log.Println("insert user error: " + err.Error())
		return
	}

	u.ID, err = result.LastInsertId()
	return
}

func SelectUser(db *sql.DB, user *models.User) (err error) {
	//TODO check if responsible is necessary
	row := db.QueryRow("SELECT * FROM Usuarios WHERE idUsuario = ?", user.ID)
	if row.Err() != nil {
		log.Println("select user error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&user.ID, &user.Person.ID, &user.CellNumber, &user.PhoneNumber, &user.Address.ID, &user.CPF, &user.Email, &user.Responsible.ID, &user.CreationDate)
	if err != nil {
		log.Println("scan user error: " + err.Error())
		return
	}

	user.CreationDate = user.CreationDate[:10]
	return
}

func SelectUsers(db *sql.DB) ([]models.User, error) {
	users := make([]models.User, 0)

	//TODO da pra melhorar isso seguindo o padrao do service
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

func DeleteUser(db *sql.DB, u *models.User) (err error) {
	//TODO
	return nil
}

func UpdateUser(db *sql.DB, u *models.User) (err error) {
	//TODO
	return nil
}
