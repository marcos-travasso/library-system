package structs

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int     `json:"id"`
	Year   int     `json:"year"`
	Pages  int     `json:"pages"`
	Title  string  `json:"title"`
	Author Author  `json:"author"`
	Genre  []Genre `json:"genre"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (g Genre) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if g.Name == "" {
			return "", errors.New("genre has no name")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Generos(nome) values (\"%s\")", g.Name)
	case "UPDATE":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		if g.Name == "" {
			return "", errors.New("genre has no name")
		}
		sqlStatement += fmt.Sprintf("UPDATE Generos SET nome=\"%s\" WHERE idGenero = \"%d\"", g.Name, g.ID)
	case "DELETE":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Generos WHERE idGenero = \"%d\"", g.ID)
	case "SELECT":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Generos WHERE idGenero = \"%d\"", g.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}
