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

func (dbDir Database) SelectLendings(e entityID) ([]structs.Lending, error) {
	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	rows, err := db.Query("SELECT COUNT(*) FROM emprestimos WHERE " + e.GetIDString())
	if err != nil {
		log.Printf("Fail to count: %s", err)
		return nil, err
	}

	lendingCount := 0
	for rows.Next() {
		err = rows.Scan(&lendingCount)
		if err != nil {
			log.Printf("Fail to receive count: %s", err)
			return nil, err
		}
	}

	lendings := make([]structs.Lending, lendingCount, lendingCount)

	rows, err = db.Query("SELECT idEmprestimo, livro, usuario, dataDoPedido from emprestimos where " + e.GetIDString())
	if err != nil {
		log.Printf("Fail to query lendings: %s", err)
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&lendings[i].ID, &lendings[i].Book.ID, &lendings[i].User.ID, &lendings[i].LendDay)
		if err != nil {
			log.Printf("Fail to receive lendings: %s", err)
			return nil, err
		}
	}

	for _, lending := range lendings {
		lending.Book, err = dbDir.SelectBook(lending.Book)
		if err != nil {
			log.Printf("Fail to query book from lending: %s", err)
			return nil, err
		}

		lending.User, err = dbDir.SelectUser(lending.User)
		if err != nil {
			log.Printf("Fail to query user from lending: %s", err)
			return nil, err
		}

		rows, err = db.Query(fmt.Sprintf("SELECT COUNT(*) FROM devolucoes WHERE emprestimo = %d", lending.ID))
		if err != nil {
			log.Printf("Fail to count: %s", err)
			return nil, err
		}

		devolutionCount := 0
		for rows.Next() {
			err = rows.Scan(&devolutionCount)
			if err != nil {
				log.Printf("Fail to receive count: %s", err)
				return nil, err
			}
		}

		rows, err = db.Query(lending.LinkSQLStatement("SELECT"))

		for rows.Next() {
			err = rows.Scan(&lending.Devolution.ID, &lending.Devolution.Date)
			if err != nil {
				log.Printf("Fail to receive devolution: %s", err)
				return nil, err
			}
		}
	}

	return lendings, nil
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
