package services

import (
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

func InsertAuthor(a *models.Author) (err error) {
	err = repositories.CheckIfAuthorExists(db, a)
	if err != nil {
		log.Println("check if author exists error: " + err.Error())
		return
	}

	if a.ID != 0 {
		return
	}

	err = repositories.InsertPerson(db, &a.Person)
	if err != nil {
		log.Println("insert author person error: " + err.Error())
		return
	}

	err = repositories.InsertAuthor(db, a)
	if err != nil {
		log.Println("insert author error: " + err.Error())
		return
	}

	return
}
