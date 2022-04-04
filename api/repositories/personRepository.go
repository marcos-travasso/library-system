package repositories

import (
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertPerson(p models.Person) (int64, error) {
	db := initializeDatabase()
	defer db.Close()

	result, err := db.Exec("INSERT INTO Pessoas(Nome, Genero, Nascimento) values (?, ?, ?)", p.Name, p.Gender, p.Birthday)
	if err != nil {
		log.Println("insert person error: " + err.Error())
		return 0, err
	}

	return result.LastInsertId()
}

func DeletePerson(p models.Person) (err error) {
	db := initializeDatabase()
	defer db.Close()

	_, err = db.Exec("DELETE FROM Pessoas WHERE idPessoa = ?", p.ID)
	if err != nil {
		log.Println("delete person error: " + err.Error())
	}

	return
}

func UpdatePerson(p models.Person) (err error) {
	db := initializeDatabase()
	defer db.Close()

	_, err = db.Exec("UPDATE Pessoas SET nome=?, genero=?, nascimento=? WHERE idPessoa = ?", p.Name, p.Gender, p.Birthday, p.ID)
	if err != nil {
		log.Println("update person error: " + err.Error())
	}

	return
}
