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

func SelectBook(db *sql.DB, b *models.Book) (models.Book, error) {
	var book models.Book

	//TODO remove this horrendous query and splice into 3 functions (selectBook, selectAuthor and selectGenres)
	row := db.QueryRow("SELECT idLivro, titulo, ano, paginas, autor, idPessoa, nome, genero, nascimento FROM (Livros INNER JOIN Autores A on Livros.autor = A.idAutor) INNER JOIN Pessoas on pessoa = Pessoas.idPessoa WHERE idLivro = ?", b.ID)
	if row.Err() != nil {
		log.Println("select book error: " + row.Err().Error())
		return book, row.Err()
	}

	err := row.Scan(&book.ID, &book.Title, &book.Year, &book.Pages, &book.Author.ID, &book.Author.Person.ID, &book.Author.Person.Name, &book.Author.Person.Gender, &book.Author.Person.Birthday)
	if err != nil {
		log.Println("scan book error: " + err.Error())
		return book, row.Err()
	}
	//TODO select genres
	return book, nil
}

func SelectBooks(db *sql.DB) ([]models.Book, error) {
	books := make([]models.Book, 0)

	rows, err := db.Query("SELECT idLivro, titulo, ano, paginas, autor, idPessoa, nome, genero, nascimento FROM (Livros INNER JOIN Autores A on Livros.autor = A.idAutor) INNER JOIN Pessoas on pessoa = Pessoas.idPessoa")
	if err != nil {
		log.Println("select books error: " + err.Error())
		return books, err
	}

	for rows.Next() {
		var newBook models.Book

		err = rows.Scan(&newBook.ID, &newBook.Title, &newBook.Year, &newBook.Pages, &newBook.Author.ID, &newBook.Author.Person.ID, &newBook.Author.Person.Name, &newBook.Author.Person.Gender, &newBook.Author.Person.Birthday)
		if err != nil {
			log.Println("scan book error: " + err.Error())
			return nil, err
		}

		books = append(books, newBook)
	}

	//TODO select genres
	return books, nil
}

func DeleteBook(db *sql.DB, b *models.Book) error {
	//TODO
	return nil
}

func UpdateBook(db *sql.DB, b *models.Book) error {
	//TODO
	return nil
}
