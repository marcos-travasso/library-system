package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertLending(db *sql.DB, l *models.Lending) (err error) {
	result, err := db.Exec("INSERT INTO emprestimos(livro, usuario, datadopedido) values (?, ?, ?)", l.Book.ID, l.User.ID, l.LendDay)
	if err != nil {
		log.Println("insert lending error: " + err.Error())
		return
	}

	l.ID, err = result.LastInsertId()
	return
}

func SelectLending(db *sql.DB, l models.Lending) (models.Lending, error) {
	//TODO
	var lend models.Lending
	return lend, nil
}

func SelectLendings(db *sql.DB) ([]models.Lending, error) {
	//TODO
	return nil, nil
}

func ReturnBook(db *sql.DB, l models.Lending) error {
	//TODO
	return nil
}

func haveLending(db *sql.DB, u models.User) (bool, error) {
	//TODO
	return false, nil
}

func isLending(db *sql.DB, b models.Book) (bool, error) {
	//TODO
	return false, nil
}
