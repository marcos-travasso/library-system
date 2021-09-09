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
