package repositories

import (
	"database/sql"
	"errors"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func CheckIfGenreExists(db *sql.DB, g *models.Genre) (err error) {
	row := db.QueryRow("SELECT idGenero from Generos where lower(nome) == ?", g.Name)
	if row.Err() != nil {
		log.Println("check if genre exists error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&g.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("scan genre id error: " + err.Error())
		return
	}

	return nil
}

func InsertGenre(db *sql.DB, g *models.Genre) (err error) {
	result, err := db.Exec("INSERT INTO Generos(nome) values (?)", g.Name)
	if err != nil {
		log.Println("insert genre error: " + err.Error())
		return err
	}

	g.ID, err = result.LastInsertId()
	return
}
