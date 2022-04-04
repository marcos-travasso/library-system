package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID           int     `json:"id"`
	Person       Person  `json:"person"`
	CellNumber   string  `json:"cellNumber"`
	PhoneNumber  string  `json:"phoneNumber"`
	CPF          string  `json:"cpf"`
	Email        string  `json:"email"`
	Address      Address `json:"address"`
	Responsible  Person  `json:"responsible"`
	CreationDate string  `json:"creationDate"`
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
		sqlStatement += fmt.Sprintf("SELECT idUsuario, celular, telefone, cpf, email, responsavel, criacao, idPessoa, nome, genero, nascimento, idEndereco, cep, cidade, bairro, rua, numero, complemento FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco) WHERE idUsuario = \"%d\"", u.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}
