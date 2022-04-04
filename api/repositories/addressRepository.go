package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertAddress(db *sql.DB, a *models.Address) (err error) {
	result, err := db.Exec("INSERT INTO Enderecos(CEP, cidade, bairro, rua, numero, complemento) values (?, ?, ?, ?, ?, ?)", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement)
	if err != nil {
		log.Println("insert address error: " + err.Error())
		return
	}

	a.ID, err = result.LastInsertId()
	return
}

func SelectAddress(db *sql.DB, a *models.Address) (err error) {
	row := db.QueryRow("SELECT * FROM Enderecos WHERE idEndereco = ?", a.ID)
	if row.Err() != nil {
		log.Println("select address error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&a.ID, &a.CEP, &a.City, &a.Neighborhood, &a.Street, &a.Number, &a.Complement)
	if err != nil {
		log.Println("scan address error: " + err.Error())
		return
	}

	return
}

func DeleteAddress(db *sql.DB, a *models.Address) (err error) {
	_, err = db.Exec("DELETE FROM Enderecos WHERE idEndereco = ?", a.ID)
	if err != nil {
		log.Println("delete address error: " + err.Error())
		return
	}

	return
}

func UpdateAddress(db *sql.DB, a *models.Address) (err error) {
	_, err = db.Exec("UPDATE Enderecos SET CEP=?, cidade=?, bairro=?, rua=?, numero=?, complemento=? WHERE idEndereco = ?", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement, a.ID)
	if err != nil {
		log.Println("update address error: " + err.Error())
		return
	}

	return
}
