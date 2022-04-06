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

	if len(user.CreationDate) > 10 {
		user.CreationDate = user.CreationDate[:10]
	}
	return
}

func SelectUsers(db *sql.DB) (users []models.User, err error) {
	rows, err := db.Query("SELECT idUsuario FROM Usuarios")
	if err != nil {
		log.Println("select users error: " + err.Error())
		return
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID)
		if err != nil {
			log.Println("scan user error: " + err.Error())
			return
		}

		users = append(users, user)
	}

	return
}

func DeleteUser(db *sql.DB, u *models.User) (err error) {
	//TODO
	return nil
}

//func UpdateUser(db *sql.DB, u *models.User) (err error) {
//	//TODO
//	return nil
//}
