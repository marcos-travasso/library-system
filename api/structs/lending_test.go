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
				ID:      1,
				User:    User{ID: 2},
				Book:    Book{ID: 3},
				LendDay: "01/01/2000",
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "02/01/200",
					},
				},
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
				ID:      0,
				User:    User{ID: 2},
				Book:    Book{ID: 3},
				LendDay: "01/01/2000",
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "02/01/200",
					},
				},
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
				ID:      1,
				User:    User{ID: 0},
				Book:    Book{ID: 3},
				LendDay: "01/01/2000",
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "03/01/200",
					},
				},
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
				ID:      1,
				User:    User{ID: 2},
				Book:    Book{ID: 0},
				LendDay: "01/01/2000",
				Devolution: []Devolution{
					{
						ID:   2,
						Date: "02/01/200",
					},
				},
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
		{name: "Lending without devolution",
			args: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "UPDATE Emprestimos SET livro=\"3\" usuario=\"2\" devolvido=\"0\" WHERE idEmprestimo = \"1\"",
				"DELETE": "DELETE FROM Emprestimos WHERE idEmprestimo = \"1\"",
				"SELECT": "SELECT * FROM Emprestimos WHERE idEmprestimo = \"1\"",
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

func TestLending_InsertDevolutionStatement(t *testing.T) {
	tests := []struct {
		name   string
		arg1   Lending
		wanted string
	}{
		{name: "One devolution",
			arg1: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "01/01/2000",
					},
				},
			},
			wanted: "INSERT INTO devolucoes(emprestimo, datadedevolucao) values (\"1\", \"01/01/2000\")",
		},
		{name: "Empty lending id",
			arg1: Lending{
				ID:       0,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "01/01/2000",
					},
				},
			},
			wanted: "",
		},
		{name: "Empty devolution date",
			arg1: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
				Devolution: []Devolution{
					{
						ID:   1,
						Date: "",
					},
				},
			},
			wanted: "",
		},
		{name: "Empty fields",
			arg1: Lending{
				ID:       0,
				User:     User{ID: 0},
				Book:     Book{ID: 0},
				LendDay:  "",
				Returned: 0,
				Devolution: []Devolution{
					{
						ID:   0,
						Date: "",
					},
				},
			},
			wanted: "",
		},
		{name: "Empty struct",
			wanted: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.arg1.LinkSQLStatement("INSERT"); got != tt.wanted {
				t.Errorf("%s statement got = %v, expect = %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestLending_SelectDevolutionStatement(t *testing.T) {
	tests := []struct {
		name   string
		arg1   Lending
		arg2   Devolution
		wanted string
	}{
		{name: "Select devolutions",
			arg1: Lending{
				ID:       1,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wanted: "SELECT * FROM devolucoes WHERE emprestimo = \"1\"",
		},
		{name: "Empty lending id",
			arg1: Lending{
				ID:       0,
				User:     User{ID: 2},
				Book:     Book{ID: 3},
				LendDay:  "01/01/2000",
				Returned: 0,
			},
			wanted: "",
		},
		{name: "Empty fields",
			arg1: Lending{
				ID:       0,
				User:     User{ID: 0},
				Book:     Book{ID: 0},
				LendDay:  "",
				Returned: 0,
			},
			wanted: "",
		},
		{name: "Empty struct",
			wanted: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.arg1.LinkSQLStatement("SELECT"); got != tt.wanted {
				t.Errorf("%s statement got = %v, expect = %v", tt.name, got, tt.wanted)
			}
		})
	}
}
