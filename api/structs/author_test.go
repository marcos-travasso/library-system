package structs

import "testing"

func TestAuthor_AuthorStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Author
		wantedStatements map[string]string
	}{
		{name: "One author",
			args: Author{
				ID: 1,
				Person: Person{
					ID:   1,
					Name: "Clarice Lispector",
				},
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Autores(pessoa) values (\"1\")",
				"UPDATE": "",
				"DELETE": "DELETE FROM Autores WHERE idAutor = \"1\"",
				"SELECT": "SELECT idAutor, nome FROM Autores INNER JOIN Pessoas ON Pessoas.idPessoa = Autores.pessoa WHERE nome = \"Clarice Lispector\"",
				"TEST":   "",
			},
		},
		{name: "Empty author ID",
			args: Author{
				ID:     0,
				Person: Person{ID: 1},
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Autores(pessoa) values (\"1\")",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty person ID",
			args: Author{
				ID:     1,
				Person: Person{ID: 0},
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Autores WHERE idAutor = \"1\"",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: Author{
				ID:     0,
				Person: Person{ID: 0},
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"EXIST":  "",
				"TEST":   "",
			},
		},
		{name: "Empty struct",
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"EXIST":  "",
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
