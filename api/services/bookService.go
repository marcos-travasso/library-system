package services

import (
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

func InsertBook(b *models.Book) (err error) {
	err = InsertAuthor(&b.Author)
	if err != nil {
		log.Println("insert author error: " + err.Error())
		return
	}

	err = InsertGenre(&b.Genre)
	if err != nil {
		log.Println("insert genre error: " + err.Error())
		return
	}

	err = repositories.InsertBook(db, b)
	if err != nil {
		log.Println("insert book error: " + err.Error())
		return
	}

	err = repositories.LinkGenre(db, b)
	if err != nil {
		log.Println("linking book and genre error: " + err.Error())
		return
	}

	return
}
