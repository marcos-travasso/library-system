package api

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

type Author struct {
	ID     int    `json:"id"`
	Person Person `json:"person"`
}

type User struct {
	ID          int     `json:"id"`
	Person      Person  `json:"person"`
	CellNumber  string  `json:"cellNumber"`
	PhoneNumber string  `json:"phoneNumber"`
	CPF         string  `json:"cpf"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
}

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
}

type Lending struct {
	ID         int          `json:"id"`
	User       User         `json:"user"`
	Book       Book         `json:"book"`
	LendDay    string       `json:"lendDay"`
	Returned   bool         `json:"returned"`
	Devolution []Devolution `json:"devolution"`
}

type Devolution struct {
	Date string `json:"date"`
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
