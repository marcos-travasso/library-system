package structs

import (
	"errors"
	"fmt"
)

type Lending struct {
	ID         int          `json:"id"`
	User       User         `json:"user"`
	Book       Book         `json:"book"`
	LendDay    string       `json:"lendDay"`
	Returned   int          `json:"returned"`
	Devolution []Devolution `json:"devolution"`
}

type Devolution struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

func (l Lending) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if l.Book.ID == 0 {
			return "", errors.New("book has no ID")
		}
		if l.User.ID == 0 {
			return "", errors.New("user has no ID")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO emprestimos(livro, usuario, datadopedido) values (\"%d\", \"%d\", \"%s\")", l.Book.ID, l.User.ID, l.LendDay)
	case "UPDATE":
		if l.ID == 0 {
			return "", errors.New("lending has no ID")
		}
		if l.Book.ID == 0 {
			return "", errors.New("book has no ID")
		}
		if l.User.ID == 0 {
			return "", errors.New("user has no ID")
		}
		sqlStatement += fmt.Sprintf("UPDATE Emprestimos SET livro=\"%d\" usuario=\"%d\" devolvido=\"%d\" WHERE idEmprestimo = \"%d\"", l.Book.ID, l.User.ID, l.Returned, l.ID)
	case "DELETE":
		if l.ID == 0 {
			return "", errors.New("lending has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Emprestimos WHERE idEmprestimo = \"%d\"", l.ID)
	case "SELECT":
		if l.ID == 0 {
			return "", errors.New("lending has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Emprestimos WHERE idEmprestimo = \"%d\"", l.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}

func (l Lending) InsertDevolutionStatement(d Devolution) (string, error) {
	sqlStatement := ""
	if l.ID == 0 {
		return "", errors.New("lending has no id")
	}
	if d.Date == "" {
		return "", errors.New("devolution has no date")
	}
	sqlStatement += fmt.Sprintf("INSERT INTO devolucoes(emprestimo, datadedevolucao) values (\"%d\", \"%s\")", l.ID, d.Date)
	return sqlStatement, nil
}

func (l Lending) SelectDevolutionStatement() (string, error) {
	sqlStatement := ""
	if l.ID == 0 {
		return "", errors.New("lending has no id")
	}
	sqlStatement += fmt.Sprintf("SELECT * FROM devolucoes WHERE emprestimo = \"%d\"", l.ID)
	return sqlStatement, nil
}
