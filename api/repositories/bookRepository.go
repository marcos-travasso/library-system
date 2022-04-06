package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertBook(db *sql.DB, b *models.Book) (err error) {
	result, err := db.Exec("INSERT INTO Livros(titulo, ano, autor, paginas) values (?, ?, ?, ?)", b.Title, b.Year, b.Author.ID, b.Pages)
	if err != nil {
		log.Println("insert book error: " + err.Error())
		return
	}

	b.ID, err = result.LastInsertId()
	return
}

func SelectBook(db *sql.DB, b *models.Book) (err error) {
	row := db.QueryRow("SELECT * FROM Livros WHERE idLivro = ?", b.ID)
	if row.Err() != nil {
		log.Println("select book error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&b.ID, &b.Title, &b.Year, &b.Author.ID, &b.Pages)
	if err != nil {
		log.Println("scan book error: " + err.Error())
		return
	}

	//TODO change this (add column genre and remove bookgenre table)
	row = db.QueryRow("SELECT * FROM generos_dos_livros where livro = ?", b.ID)
	if row.Err() != nil {
		log.Println("select book genre error: " + row.Err().Error())
		return row.Err()
	}

	err = row.Scan(&b.ID, &b.Genre.ID)
	if err != nil {
		log.Println("scan book genre error: " + err.Error())
		return
	}

	return
}

func SelectBooks(db *sql.DB) (books []models.Book, err error) {
	rows, err := db.Query("SELECT idLivro FROM Livros")
	if err != nil {
		log.Println("select books error: " + err.Error())
		return
	}

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID)
		if err != nil {
			log.Println("scan book error: " + err.Error())
			return
		}

		books = append(books, book)
	}

	return
}

func DeleteBook(db *sql.DB, b *models.Book) error {
	//TODO
	return nil
}

//func UpdateBook(db *sql.DB, b *models.Book) error {
//	//TODO
//	return nil
//}
