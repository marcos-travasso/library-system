package structs

import (
	"testing"
)

func TestUser_UserStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             User
		wantedStatements map[string]string
	}{
		{name: "One user",
			args: User{
				ID: 1,
				Person: Person{
					ID: 2,
				},
				CellNumber:  "12345678910",
				PhoneNumber: "9876543210",
				CPF:         "12345678910",
				Email:       "testmail@mail.test",
				Address:     Address{ID: 3},
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Usuarios(pessoa, celular, telefone, endereco, cpf, email, criacao) values (\"2\", \"12345678910\", \"9876543210\", \"3\", \"12345678910\", \"testmail@mail.test\", date('now'))",
				"UPDATE": "UPDATE Usuarios SET pessoa=\"2\", celular=\"12345678910\", telefone=\"9876543210\", endereco=\"3\", cpf=\"12345678910\", email=\"testmail@mail.test\" WHERE idUsuario = \"1\"",
				"DELETE": "DELETE FROM Usuarios WHERE idUsuario = \"1\"",
				"SELECT": "SELECT * FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco) WHERE idUsuario = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty person ID",
			args: User{
				ID: 1,
				Person: Person{
					ID: 0,
				},
				CellNumber:  "12345678910",
				PhoneNumber: "9876543210",
				CPF:         "12345678910",
				Email:       "testmail@mail.test",
				Address:     Address{ID: 3},
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Usuarios WHERE idUsuario = \"1\"",
				"SELECT": "SELECT * FROM ((Usuarios INNER JOIN Pessoas ON Pessoas.idPessoa = Usuarios.pessoa) INNER JOIN Enderecos ON Usuarios.endereco = Enderecos.idEndereco) WHERE idUsuario = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty user ID",
			args: User{
				ID: 0,
				Person: Person{
					ID: 2,
				},
				CellNumber:  "12345678910",
				PhoneNumber: "9876543210",
				CPF:         "12345678910",
				Email:       "testmail@mail.test",
				Address:     Address{ID: 3},
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Usuarios(pessoa, celular, telefone, endereco, cpf, email, criacao) values (\"2\", \"12345678910\", \"9876543210\", \"3\", \"12345678910\", \"testmail@mail.test\", date('now'))",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: User{
				ID: 0,
				Person: Person{
					ID: 0,
				},
				CellNumber:  "",
				PhoneNumber: "",
				CPF:         "",
				Email:       "",
				Address:     Address{},
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty struct",
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for statementType, wanted := range tt.wantedStatements {
				if got, _ := tt.args.SQLStatement(statementType); got != wanted {
					t.Errorf("%s statement got = %v, expect = %v", statementType, got, wanted)
				}
			}
		})
	}
}

func TestUser_AddressStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Address
		wantedStatements map[string]string
	}{
		{name: "One address",
			args: Address{
				ID:           1,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Centro",
				Street:       "XV",
				Complement:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Enderecos(CEP, cidade, bairro, rua, numero, complemento) values (\"12345678\", \"Curitiba\", \"Centro\", \"XV\", \"123\", \"\")",
				"UPDATE": "UPDATE Enderecos SET CEP=\"12345678\", cidade=\"Curitiba\", bairro=\"Centro\", rua=\"XV\", numero=\"123\", complemento=\"\" WHERE idEndereco = \"1\"",
				"DELETE": "DELETE FROM Enderecos WHERE idEndereco = \"1\"",
				"SELECT": "SELECT * FROM Enderecos WHERE idEndereco = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty id",
			args: Address{
				ID:           0,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Centro",
				Street:       "XV",
				Complement:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Enderecos(CEP, cidade, bairro, rua, numero, complemento) values (\"12345678\", \"Curitiba\", \"Centro\", \"XV\", \"123\", \"\")",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty arguments",
			args: Address{
				ID:           1,
				Number:       0,
				CEP:          "",
				City:         "",
				Neighborhood: "",
				Street:       "",
				Complement:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Enderecos WHERE idEndereco = \"1\"",
				"SELECT": "SELECT * FROM Enderecos WHERE idEndereco = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: Address{
				ID:           0,
				Number:       0,
				CEP:          "",
				City:         "",
				Neighborhood: "",
				Street:       "",
				Complement:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty struct",
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for statementType, wanted := range tt.wantedStatements {
				if got, _ := tt.args.SQLStatement(statementType); got != wanted {
					t.Errorf("%s statement got = %v, expect = %v", statementType, got, wanted)
				}
			}
		})
	}
}
