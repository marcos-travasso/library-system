package structs

import "testing"

func TestBook_GenreStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Genre
		wantedStatements map[string]string
	}{
		{name: "One genre",
			args: Genre{
				ID:   1,
				Name: "Science Fiction",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Generos(nome) values (\"Science Fiction\")",
				"UPDATE": "UPDATE Generos SET nome=\"Science Fiction\" WHERE idGenero = \"1\"",
				"DELETE": "DELETE FROM Generos WHERE idGenero = \"1\"",
				"SELECT": "SELECT * FROM Generos WHERE idGenero = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty name",
			args: Genre{
				ID:   2,
				Name: "",
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Generos WHERE idGenero = \"2\"",
				"SELECT": "SELECT * FROM Generos WHERE idGenero = \"2\"",
				"TEST":   "",
			},
		},
		{name: "Empty ID",
			args: Genre{
				ID:   0,
				Name: "Science Fiction",
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Generos(nome) values (\"Science Fiction\")",
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
