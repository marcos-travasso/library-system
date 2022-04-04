package models

import (
	"errors"
	"fmt"
)

type Address struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
}

func (a Address) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if a.CEP == "" && a.City == "" && a.Neighborhood == "" && a.Street == "" {
			return "", errors.New("address has no arguments")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Enderecos(CEP, cidade, bairro, rua, numero, complemento) values (\"%s\", \"%s\", \"%s\", \"%s\", \"%d\", \"%s\")", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement)
	case "UPDATE":
		if a.CEP == "" && a.City == "" && a.Neighborhood == "" && a.Street == "" {
			return "", errors.New("address has no arguments")
		}
		if a.ID == 0 {
			return "", errors.New("address has no ID")
		}
		sqlStatement += fmt.Sprintf("UPDATE Enderecos SET CEP=\"%s\", cidade=\"%s\", bairro=\"%s\", rua=\"%s\", numero=\"%d\", complemento=\"%s\" WHERE idEndereco = \"%d\"", a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement, a.ID)
	case "DELETE":
		if a.ID == 0 {
			return "", errors.New("address has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Enderecos WHERE idEndereco = \"%d\"", a.ID)
	case "SELECT":
		if a.ID == 0 {
			return "", errors.New("address has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Enderecos WHERE idEndereco = \"%d\"", a.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}
