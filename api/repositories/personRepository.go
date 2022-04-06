package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertPerson(db *sql.DB, p *models.Person) (err error) {
	result, err := db.Exec("INSERT INTO Pessoas(Nome, Genero, Nascimento) values (?, ?, ?)", p.Name, p.Gender, p.Birthday)
	if err != nil {
		log.Println("insert person error: " + err.Error())
		return
	}

	p.ID, err = result.LastInsertId()
	return
}

func SelectPerson(db *sql.DB, p *models.Person) (err error) {
	row := db.QueryRow("SELECT * FROM Pessoas WHERE idPessoa = ?", p.ID)
	if row.Err() != nil {
		log.Println("select person error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&p.ID, &p.Name, &p.Gender, &p.Birthday)
	if err != nil {
		log.Println("scan person error: " + err.Error())
		return
	}

	return
}

func DeletePerson(db *sql.DB, p *models.Person) (err error) {
	_, err = db.Exec("DELETE FROM Pessoas WHERE idPessoa = ?", p.ID)
	if err != nil {
		log.Println("delete person error: " + err.Error())
		return
	}

	return
}

//func UpdatePerson(db *sql.DB, p *models.Person) (err error) {
//	_, err = db.Exec("UPDATE Pessoas SET nome=?, genero=?, nascimento=? WHERE idPessoa = ?", p.Name, p.Gender, p.Birthday, p.ID)
//	if err != nil {
//		log.Println("update person error: " + err.Error())
//		return
//	}
//
//	return
//}
