package repositories

import (
	"github.com/marcos-travasso/library-system/models"
)

func InsertLending(l models.Lending) (int, error) {
	//TODO
	return 0, nil
}

func SelectLending(l models.Lending) (models.Lending, error) {
	//TODO
	var lend models.Lending
	return lend, nil
}

func SelectLendings() ([]models.Lending, error) {
	//TODO
	return nil, nil
}

func ReturnBook(l models.Lending) error {
	//TODO
	return nil
}

func haveLending(u models.User) (bool, error) {
	//TODO
	return false, nil
}

func isLending(b models.Book) (bool, error) {
	//TODO
	return false, nil
}
