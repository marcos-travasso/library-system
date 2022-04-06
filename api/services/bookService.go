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

func SelectBook(b *models.Book) (err error) {
	err = repositories.SelectBook(db, b)
	if err != nil {
		log.Println("select book error: " + err.Error())
		return
	}

	err = SelectGenre(&b.Genre)
	if err != nil {
		log.Println("select book genre error: " + err.Error())
		return
	}

	err = SelectAuthor(&b.Author)
	if err != nil {
		log.Println("select book author error: " + err.Error())
		return
	}

	return
}

func SelectBooks() (books []models.Book, err error) {
	books, err = repositories.SelectBooks(db)
	if err != nil {
		log.Println("select books error: " + err.Error())
		return
	}

	for i := range books {
		err = SelectBook(&books[i])
		if err != nil {
			log.Println("select book error: " + err.Error())
			return
		}
	}

	return
}
