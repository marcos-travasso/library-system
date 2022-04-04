package repositories

import (
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func CheckIfGenreExists(g models.Genre) (genreId int, err error) {
	db := initializeDatabase()
	defer db.Close()

	row := db.QueryRow("select idGenero from Generos where lower(nome) == ?", g.Name)
	if row.Err() != nil {
		log.Println("check if genre exists error: " + row.Err().Error())
		return 0, row.Err()
	}

	err = row.Scan(&genreId)
	if err != nil {
		log.Println("scan genre id error: " + err.Error())
	}
	//TODO verify if it returns int64 or int
	return
}

func InsertGenre(g models.Genre) (int, error) {
	db := initializeDatabase()
	defer db.Close()

	result, err := db.Exec("INSERT INTO Generos(nome) values (?)", g.Name)
	if err != nil {
		log.Println("insert genre error: " + err.Error())
		return 0, err
	}

	genreId, _ := result.LastInsertId()
	g.ID = int(genreId)

	return g.ID, nil
}
