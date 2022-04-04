package models

import (
	"errors"
	"fmt"
)

type Lending struct {
	ID         int        `json:"id"`
	User       User       `json:"user"`
	Book       Book       `json:"book"`
	LendDay    string     `json:"lendDay"`
	Returned   int        `json:"returned"`
	Devolution Devolution `json:"devolution"`
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
		if l.Devolution.Date == "" {
			return "", errors.New("lending has no devolution")
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

func (l Lending) LinkSQLStatement(statementType string) (string, error) {
	if l.ID == 0 {
		return "", errors.New("lending has no id")
	}

	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if l.Devolution.Date == "" {
			return "", errors.New("devolution has no date")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO devolucoes(emprestimo, datadedevolucao) values (\"%d\", \"%s\")", l.ID, l.Devolution.Date)
	case "SELECT":
		sqlStatement += fmt.Sprintf("SELECT idDevolucao, dataDeDevolucao FROM devolucoes WHERE emprestimo = \"%d\"", l.ID)
	default:
		return "", errors.New("invalid statement type")
	}

	return sqlStatement, nil
}
