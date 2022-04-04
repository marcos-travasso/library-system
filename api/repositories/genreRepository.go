package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func CheckIfGenreExists(db *sql.DB, g models.Genre) (genreId int, err error) {
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

func InsertGenre(db *sql.DB, g models.Genre) (int64, error) {
	result, err := db.Exec("INSERT INTO Generos(nome) values (?)", g.Name)
	if err != nil {
		log.Println("insert genre error: " + err.Error())
		return 0, err
	}

	genreId, _ := result.LastInsertId()
	g.ID = genreId

	return g.ID, nil
}
