package structs

import (
	"errors"
	"fmt"
)

type Author struct {
	ID     int    `json:"id"`
	Person Person `json:"person"`
}

func (a Author) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if a.Person.ID == 0 {
			return "", errors.New("author has no person ID")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Autores(pessoa) values (\"%d\")", a.Person.ID)
	case "DELETE":
		if a.ID == 0 {
			return "", errors.New("author has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Autores WHERE idAutor = \"%d\"", a.ID)
	case "SELECT":
		if a.ID == 0 {
			return "", errors.New("author has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Autores WHERE idAutor = \"%d\"", a.ID)
	}
	return sqlStatement, nil
}
