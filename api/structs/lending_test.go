package structs

import "testing"

func TestLending_LendingStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Lending
		wantedStatements map[string]string
	}{
		{name: "One lending",
			args: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO emprestimos(livro, usuario, datadopedido) values (\"3\", \"2\", \"01/01/2000\")",
				"UPDATE": "UPDATE Emprestimos SET livro=\"3\" usuario=\"2\" devolvido=\"0\" WHERE idEmprestimo = \"1\"",
				"DELETE": "DELETE FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"SELECT": "SELECT * FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty id",
			args: Lending{
				ID:       0,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO emprestimos(livro, usuario, datadopedido) values (\"3\", \"2\", \"01/01/2000\")",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty user id",
			args: Lending{
				ID:       1,
				User:     User{ID: 0},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"SELECT": "SELECT * FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty book id",
			args: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 0},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"SELECT": "SELECT * FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: Lending{
				ID:       0,
				User:     User{ID: 0},
				Book:     Book{ID: 0},
				LendDay:  "",
				Returned: 0,
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
