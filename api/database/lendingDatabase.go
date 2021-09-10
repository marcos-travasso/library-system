package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
)

func (dbDir Database) InsertLending(l structs.Lending) (int, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	haveLending, err := dbDir.haveLending(l.User)
	if err != nil {
		return 0, err
	}
	if haveLending {
		return 0, errors.New("user already have lending")
	}

	isLending, err := dbDir.isLending(l.Book)
	if err != nil {
		return 0, err
	}
	if isLending {
		return 0, errors.New("book is already lending")
	}

	err = sendStatement(l, "INSERT", db)
	if err != nil {
		return 0, err
	}

	l.ID, err = dbDir.getLastID("emprestimos", "idEmprestimo")
	if err != nil {
		return 0, err
	}

	err = sendLinkStatement(l, "INSERT", db)
	if err != nil {
		return 0, err
	}

	return l.ID, nil
}

func (dbDir *Database) haveLending(u structs.User) (bool, error) {
	var db = initializeDatabase(*dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	query := fmt.Sprintf("SELECT devolvido from emprestimos WHERE usuario = %d ORDER BY dataDoPedido DESC LIMIT 1", u.ID)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	haveLendingInt := 0
	haveLending := false
	for rows.Next() {
		err = rows.Scan(&haveLendingInt)
		if err != nil {
			return false, err
		}

		haveLending = haveLendingInt == 0
	}

	return haveLending, nil
}

func (dbDir *Database) isLending(b structs.Book) (bool, error) {
	var db = initializeDatabase(*dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	query := fmt.Sprintf("SELECT devolvido from emprestimos WHERE livro = %d ORDER BY dataDoPedido DESC LIMIT 1", b.ID)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	isLendingInt := 0
	isLending := false
	for rows.Next() {
		err = rows.Scan(&isLendingInt)
		if err != nil {
			return false, err
		}

		isLending = isLendingInt == 0
	}

	return isLending, nil
}
