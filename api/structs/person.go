package structs

import (
	"errors"
	"fmt"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func (p Person) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if p.Name == "" {
			return "", errors.New("person has no name")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"%s\", \"%s\", \"%s\")", p.Name, p.Gender, p.Birthday)
	case "UPDATE":
		if p.ID == 0 {
			return "", errors.New("person has no ID")
		}
		sqlStatement += fmt.Sprintf("UPDATE Pessoas SET nome=\"%s\", genero=\"%s\", nascimento=\"%s\" WHERE idPessoa = \"%d\"", p.Name, p.Gender, p.Birthday, p.ID)
	case "DELETE":
		if p.ID == 0 {
			return "", errors.New("person has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Pessoas WHERE idPessoa = \"%d\"", p.ID)
	case "SELECT":
		if p.ID == 0 {
			return "", errors.New("person has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Pessoas WHERE idPessoa = \"%d\"", p.ID)
	}
	return sqlStatement, nil
}
