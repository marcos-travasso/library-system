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
	return l.ID, err
}

func (dbDir Database) SelectLending(l structs.Lending) (structs.Lending, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	rows, err := db.Query(fmt.Sprintf("SELECT livro, usuario, dataDoPedido, devolvido from emprestimos WHERE idEmprestimo = %d", l.ID))
	if err != nil {
		return l, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&l.Book.ID, &l.User.ID, &l.LendDay, &l.Returned)
		if err != nil {
			return l, err
		}
	}

	if l.Book.ID == 0 {
		return l, errors.New("no lending was found with this ID")
	}

	l.Book, err = dbDir.SelectBook(l.Book)
	if err != nil {
		return l, err
	}

	l.User, err = dbDir.SelectUser(l.User)
	if err != nil {
		return l, err
	}

	rows, err = db.Query(l.LinkSQLStatement("SELECT"))

	for rows.Next() {
		err = rows.Scan(&l.Devolution.ID, &l.Devolution.Date)
		if err != nil {
			return l, err
		}
	}

	return l, nil
}

func (dbDir Database) SelectLendings() ([]structs.Lending, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	rows, err := db.Query("SELECT COUNT(*) FROM emprestimos")
	if err != nil {
		return nil, err
	}

	lendingCount := 0
	for rows.Next() {
		err = rows.Scan(&lendingCount)
		if err != nil {
			return nil, err
		}
	}

	lendings := make([]structs.Lending, lendingCount, lendingCount)

	rows, err = db.Query("SELECT idEmprestimo, livro, usuario, dataDoPedido, devolvido from emprestimos")
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&lendings[i].ID, &lendings[i].Book.ID, &lendings[i].User.ID, &lendings[i].LendDay, &lendings[i].Returned)
		if err != nil {
			return nil, err
		}
	}

	for i := range lendings {
		lendings[i].Book, err = dbDir.SelectBook(lendings[i].Book)
		if err != nil {
			return nil, err
		}

		lendings[i].User, err = dbDir.SelectUser(lendings[i].User)
		if err != nil {
			return nil, err
		}

		rows, err = db.Query(lendings[i].LinkSQLStatement("SELECT"))

		for rows.Next() {
			err = rows.Scan(&lendings[i].Devolution.ID, &lendings[i].Devolution.Date)
			if err != nil {
				return nil, err
			}
		}
	}

	return lendings, nil
}

func (dbDir Database) ReturnBook(l structs.Lending) error {
	if l.ID == 0 {
		return errors.New("lending has no id")
	}

	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	_, err := db.Exec(fmt.Sprintf("UPDATE emprestimos SET devolvido = 1 WHERE idEmprestimo = %d", l.ID))
	return err
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
