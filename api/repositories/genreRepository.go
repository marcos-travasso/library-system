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
		return
	}

	g.ID, err = result.LastInsertId()
	return
}

func SelectGenre(db *sql.DB, g *models.Genre) (err error) {
	row := db.QueryRow("SELECT * from Generos where idGenero == ?", g.ID)
	if row.Err() != nil {
		log.Println("select genre error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&g.ID, &g.Name)
	if err != nil {
		log.Println("scan genre error: " + err.Error())
		return
	}

	return
}

func LinkGenre(db *sql.DB, b *models.Book) (err error) {
	_, err = db.Exec("INSERT INTO generos_dos_livros(livro, genero) VALUES (?, ?)", b.ID, b.Genre.ID)
	if err != nil {
		log.Println("linking book and genre error: " + err.Error())
		return
	}

	return
}
