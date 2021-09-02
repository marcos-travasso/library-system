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
				ID:     1,
				Person: Person{ID: 1},
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Autores(pessoa) values (\"1\")",
				"UPDATE": "",
				"DELETE": "DELETE FROM Autores WHERE idAutor = \"1\"",
				"SELECT": "SELECT * FROM Autores WHERE idAutor = \"1\"",
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
				"SELECT": "SELECT * FROM Autores WHERE idAutor = \"1\"",
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
