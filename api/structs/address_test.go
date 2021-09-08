package structs

import "testing"

func TestAddress_AddressStatement(t *testing.T) {
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
