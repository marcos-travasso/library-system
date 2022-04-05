package services

import (
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

func InsertGenre(g *models.Genre) (err error) {
	err = repositories.CheckIfGenreExists(db, g)
	if err != nil {
		log.Println("check if genre exists error: " + err.Error())
		return
	}

	if g.ID != 0 {
		return
	}

	err = repositories.InsertGenre(db, g)
	if err != nil {
		log.Println("insert genre error: " + err.Error())
		return
	}

	return
}

func SelectGenre(g *models.Genre) (err error) {
	err = repositories.SelectGenre(db, g)
	if err != nil {
		log.Println("select genre error: " + err.Error())
		return
	}

	return
}
