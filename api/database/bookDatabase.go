package database

import (
	"database/sql"
	"errors"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) InsertBook(b structs.Book) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	authorID, err := dbDir.InsertAuthor(b.Author)
	if err != nil {
		return 0, err
	}
	b.Author.ID = authorID

	genreID, err := dbDir.insertGenre(b.Genre)
	if err != nil {
		return 0, err
	}
	b.Genre.ID = genreID

	err = sendLinkStatement(b, "INSERT", db)
	if err != nil {
		return 0, err
	}

	err = sendStatement(b, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return dbDir.getLastID("Livros", "idLivro")
}

func (dbDir Database) SelectBook(b structs.Book) (structs.Book, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	book := structs.Book{}

	rows, err := db.Query(b.SQLStatement("SELECT"))
	if err != nil {
		log.Printf("Fail to query book id: %s", err)
		return book, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Print(err)
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Year, &book.Pages, &book.Author.ID, &book.Author.Person.ID, &book.Author.Person.Name, &book.Author.Person.Gender, &book.Author.Person.Birthday)
		if err != nil {
			log.Printf("Fail to receive book id: %s", err)
			return book, err
		}
	}

	rows, err = db.Query(b.LinkSQLStatement("SELECT"))
	if err != nil {
		log.Printf("Fail to query book genre: %s", err)
		return book, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Print(err)
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&book.Genre.ID, &book.Genre.Name)
		if err != nil {
			log.Printf("Fail to receive genre: %s", err)
			return book, err
		}
	}

	if book.ID == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (dbDir Database) SelectBooks() ([]structs.Book, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	bookCount, err := dbDir.countRows("Livros")
	if err != nil {
		log.Printf("Fail to receive book count: %s", err)
		return nil, err
	}

	books := make([]structs.Book, bookCount, bookCount)

	rows, err := db.Query("SELECT idLivro, titulo, ano, paginas, autor, idPessoa, nome, genero, nascimento FROM (Livros INNER JOIN Autores A on Livros.autor = A.idAutor) INNER JOIN Pessoas on pessoa = Pessoas.idPessoa")
	if err != nil {
		log.Printf("Fail to query books: %s", err)
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&books[i].ID, &books[i].Title, &books[i].Year, &books[i].Pages, &books[i].Author.ID, &books[i].Author.Person.ID, &books[i].Author.Person.Name, &books[i].Author.Person.Gender, &books[i].Author.Person.Birthday)
		if err != nil {
			log.Printf("Fail to receive books: %s", err)
			return nil, err
		}
	}

	for i := range books {
		rows, err = db.Query(books[i].LinkSQLStatement("SELECT"))
		if err != nil {
			log.Printf("Fail to query book genre: %s", err)
			return nil, err
		}

		for rows.Next() {
			err = rows.Scan(&books[i].Genre.ID, &books[i].Genre.Name)
			if err != nil {
				log.Printf("Fail to receive genre: %s", err)
				return nil, err
			}
		}
	}

	return books, nil
}

func (dbDir Database) DeleteBook(b structs.Book) error {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	err := sendLinkStatement(b, "DELETE", db)
	if err != nil {
		return err
	}

	err = sendStatement(b, "DELETE", db)
	if err != nil {
		return err
	}

	return nil
}
