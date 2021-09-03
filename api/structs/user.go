package structs

import (
	"errors"
	"fmt"
)

type User struct {
	ID          int     `json:"id"`
	Person      Person  `json:"person"`
	CellNumber  string  `json:"cellNumber"`
	PhoneNumber string  `json:"phoneNumber"`
	CPF         string  `json:"cpf"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
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

func (u User) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if u.Person.ID == 0 {
			return "", errors.New("user has no person ID")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Usuarios(pessoa, celular, telefone, endereco, cpf, email, criacao) values (\"%d\", \"%s\", \"%s\", \"%d\", \"%s\", \"%s\", date('now'))", u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email)
	case "UPDATE":
		if u.ID == 0 {
			return "", errors.New("user has no ID")
		}
		if u.Person.ID == 0 {
			return "", errors.New("user has no person ID")
		}
		sqlStatement += fmt.Sprintf("UPDATE Usuarios SET pessoa=\"%d\", celular=\"%s\", telefone=\"%s\", endereco=\"%d\", cpf=\"%s\", email=\"%s\" WHERE idUsuario = \"%d\"", u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email, u.ID)
	case "DELETE":
		if u.ID == 0 {
			return "", errors.New("user has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Usuarios WHERE idUsuario = \"%d\"", u.ID)
	case "SELECT":
		if u.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Usuarios WHERE idUsuario = \"%d\"", u.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
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
