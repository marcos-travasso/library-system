package services

import (
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

func InsertLending(l *models.Lending) (err error) {
	if err = repositories.IsLending(db, l); err != nil {
		return
	}

	err = repositories.InsertLending(db, l)
	if err != nil {
		log.Println("insert lending error: " + err.Error())
		return
	}

	err = repositories.InsertDevolution(db, l)
	if err != nil {
		log.Println("insert devolution error: " + err.Error())
		return
	}

	return
}
