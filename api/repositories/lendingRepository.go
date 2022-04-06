package repositories

import (
	"database/sql"
	"errors"
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

func IsLending(db *sql.DB, l *models.Lending) (err error) {
	row := db.QueryRow("SELECT livro, usuario FROM emprestimos WHERE (livro == ? OR usuario == ?) AND devolvido == 0", l.Book.ID, l.User.ID)
	if row.Err() != nil {
		log.Println("check if lending exists error: " + row.Err().Error())
		return row.Err()
	}

	var bookId, userId int64
	err = row.Scan(&bookId, &userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("scan lending error: " + err.Error())
		return
	}

	if bookId == l.Book.ID {
		return models.ErrorAlreadyLending("book")
	} else if userId == l.User.ID {
		return models.ErrorAlreadyLending("user")
	}

	return nil
}

func ReturnBook(db *sql.DB, l *models.Lending) error {
	//TODO
	return nil
}
