package repositories

import (
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func CheckIfAuthorExists(a models.Author) (authorId int, err error) {
	db := InitializeDatabase()
	defer db.Close()

	row := db.QueryRow("select idAutor from Autores inner join Pessoas P on P.idPessoa = Autores.pessoa where lower(nome) == ?", a.Person.Name)
	if row.Err() != nil {
		log.Println("check if author exists error: " + row.Err().Error())
		return 0, row.Err()
	}

	err = row.Scan(&authorId)
	if err != nil {
		log.Println("scan author id error: " + err.Error())
	}
	//TODO verify if it returns int64 or int
	return
}

//TODO call CheckIfAuthorExists before InsertAuthor in the author service file
func InsertAuthor(a models.Author) (int64, error) {
	db := InitializeDatabase()
	defer db.Close()

	result, err := db.Exec("INSERT INTO Autores(pessoa) values (?)", a.Person.ID)
	if err != nil {
		log.Println("insert author error: " + err.Error())
		return 0, err
	}

	//TODO return a models.Author struct instead of int
	return result.LastInsertId()
}
