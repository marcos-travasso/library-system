package repositories

import (
	"database/sql"
	"errors"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func CheckIfAuthorExists(db *sql.DB, a *models.Author) (err error) {
	//Maybe this inner join is useful to avoid searching for a user with the same name as the wanted author
	row := db.QueryRow("SELECT idAutor, pessoa from Autores inner join Pessoas P on P.idPessoa = Autores.pessoa where lower(nome) == ?", a.Person.Name)
	if row.Err() != nil {
		log.Println("check if author exists error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&a.ID, &a.Person.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("scan author id error: " + err.Error())
		return
	}

	return nil
}

func InsertAuthor(db *sql.DB, a *models.Author) (err error) {
	result, err := db.Exec("INSERT INTO Autores(pessoa) values (?)", a.Person.ID)
	if err != nil {
		log.Println("insert author error: " + err.Error())
		return
	}

	a.ID, err = result.LastInsertId()
	return
}

func SelectAuthor(db *sql.DB, a *models.Author) (err error) {
	row := db.QueryRow("SELECT * from Autores where idAutor == ?", a.ID)
	if row.Err() != nil {
		log.Println("select author error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&a.ID, &a.Person.ID)
	if err != nil {
		log.Println("scan author error: " + err.Error())
		return
	}

	return
}
