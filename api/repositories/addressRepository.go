package repositories

import (
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertAddress(a models.Address) (int64, error) {
	db := initializeDatabase()
	defer db.Close()

	result, err := db.Exec("INSERT INTO Enderecos(CEP, cidade, bairro, rua, numero, complemento) values (?, ?, ?, ?, ?, ?)", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement)
	if err != nil {
		log.Println("insert address error: " + err.Error())
		return 0, err
	}

	return result.LastInsertId()
}

func DeleteAddress(a models.Address) (err error) {
	db := initializeDatabase()
	defer db.Close()

	_, err = db.Exec("DELETE FROM Enderecos WHERE idEndereco = ?", a.ID)
	if err != nil {
		log.Println("delete address error: " + err.Error())
	}

	return
}

func UpdateAddress(a models.Address) (err error) {
	db := initializeDatabase()
	defer db.Close()

	_, err = db.Exec("UPDATE Enderecos SET CEP=?, cidade=?, bairro=?, rua=?, numero=?, complemento=? WHERE idEndereco = ?", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement, a.ID)
	if err != nil {
		log.Println("update address error: " + err.Error())
	}

	return
}
