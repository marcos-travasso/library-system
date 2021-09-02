package structs

import "testing"

func TestPerson_PersonStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Person
		wantedStatements map[string]string
	}{
		{name: "One person",
			args: Person{
				ID:       1,
				Name:     "Marcos",
				Birthday: "01/01/2000",
				Gender:   "M",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"01/01/2000\")",
				"UPDATE": "UPDATE Pessoas SET nome=\"Marcos\", genero=\"M\", nascimento=\"01/01/2000\" WHERE idPessoa = \"1\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"1\"",
				"SELECT": "SELECT * FROM Pessoas WHERE idPessoa = \"1\"",
			},
		},
		{name: "Empty person",
			args: Person{
				ID:       2,
				Name:     "",
				Birthday: "",
				Gender:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"2\"",
				"SELECT": "SELECT * FROM Pessoas WHERE idPessoa = \"2\"",
				"TEST":   "",
			},
		},
		{name: "Empty name",
			args: Person{
				ID:       3,
				Name:     "",
				Birthday: "01/01/2000",
				Gender:   "F",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"3\"",
				"SELECT": "SELECT * FROM Pessoas WHERE idPessoa = \"3\"",
				"TEST":   "",
			},
		},
		{name: "Empty birthday",
			args: Person{
				ID:       4,
				Name:     "Marcos",
				Birthday: "",
				Gender:   "M",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"\")",
				"UPDATE": "UPDATE Pessoas SET nome=\"Marcos\", genero=\"M\", nascimento=\"\" WHERE idPessoa = \"4\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"4\"",
				"SELECT": "SELECT * FROM Pessoas WHERE idPessoa = \"4\"",
				"TEST":   "",
			},
		},
		{name: "Empty ID",
			args: Person{
				ID:       0,
				Name:     "Marcos",
				Birthday: "01/01/2000",
				Gender:   "M",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"01/01/2000\")",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: Person{
				ID:       0,
				Name:     "",
				Birthday: "",
				Gender:   "",
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
