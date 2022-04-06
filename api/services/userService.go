package services

import (
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/repositories"
	"log"
)

func InsertUser(u *models.User) (err error) {
	err = repositories.InsertPerson(db, &u.Person)
	if err != nil {
		log.Println("insert person error: " + err.Error())
		return
	}

	err = repositories.InsertAddress(db, &u.Address)
	if err != nil {
		log.Println("insert address error: " + err.Error())
		return
	}

	err = repositories.InsertUser(db, u)
	if err != nil {
		log.Println("insert User error: " + err.Error())
		return
	}

	return
}

func SelectUser(u *models.User) (err error) {
	err = repositories.SelectUser(db, u)
	if err != nil {
		log.Println("select User error: " + err.Error())
		return
	}

	err = repositories.SelectAddress(db, &u.Address)
	if err != nil {
		log.Println("select address error: " + err.Error())
		return
	}

	err = repositories.SelectPerson(db, &u.Person)
	if err != nil {
		log.Println("select person error: " + err.Error())
		return
	}

	return
}

func SelectUsers() (users []models.User, err error) {
	users, err = repositories.SelectUsers(db)
	if err != nil {
		log.Println("select users error: " + err.Error())
		return
	}

	for i := range users {
		err = SelectUser(&users[i])
		if err != nil {
			log.Println("select user error: " + err.Error())
			return
		}
	}

	return
}

func DeleteUser(u *models.User) (err error) {
	//TODO
	return nil
}

func UpdateUser(u *models.User) (err error) {
	//TODO
	return nil
}
