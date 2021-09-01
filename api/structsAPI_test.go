package api

import "testing"

func TestPerson_PersonStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Person
		wantedStatements map[string]string
		insertWant       string
		updateWant       string
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
			},
		},
		{name: "Empty person",
			args: Person{
				ID:       1,
				Name:     "",
				Birthday: "",
				Gender:   "",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "UPDATE Pessoas SET nome=\"\", genero=\"\", nascimento=\"\" WHERE idPessoa = \"1\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"1\"",
			},
		},
		{name: "Empty name",
			args: Person{
				ID:       1,
				Name:     "",
				Birthday: "01/01/2000",
				Gender:   "F",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "UPDATE Pessoas SET nome=\"\", genero=\"F\", nascimento=\"01/01/2000\" WHERE idPessoa = \"1\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"1\"",
			},
		},
		{name: "Empty birthday",
			args: Person{
				ID:       1,
				Name:     "Marcos",
				Birthday: "",
				Gender:   "M",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"\")",
				"UPDATE": "UPDATE Pessoas SET nome=\"Marcos\", genero=\"M\", nascimento=\"\" WHERE idPessoa = \"1\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"1\"",
			},
		},
		{name: "Empty birthday",
			args: Person{
				ID:       1,
				Name:     "Marcos",
				Birthday: "",
				Gender:   "M",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"\")",
				"UPDATE": "UPDATE Pessoas SET nome=\"Marcos\", genero=\"M\", nascimento=\"\" WHERE idPessoa = \"1\"",
				"DELETE": "DELETE FROM Pessoas WHERE idPessoa = \"1\"",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for statementType, wanted := range tt.wantedStatements {
				if got := tt.args.SQLStatement(statementType); got != wanted {
					t.Errorf("insertPerson() = %v, want %v", got, wanted)
				}
			}
		})
	}
}
