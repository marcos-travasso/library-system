package database

import (
	"database/sql"
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

	return book, nil
}
