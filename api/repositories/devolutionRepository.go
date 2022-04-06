package repositories

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/models"
	"log"
)

func InsertDevolution(db *sql.DB, l *models.Lending) (err error) {
	result, err := db.Exec("INSERT INTO devolucoes(emprestimo, datadedevolucao) values (?, ?)", l.ID, l.Devolution.Date)
	if err != nil {
		log.Println("insert devolution error: " + err.Error())
		return
	}

	l.Devolution.ID, err = result.LastInsertId()
	return
}
