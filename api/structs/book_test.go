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

func TestBook_BookStatement(t *testing.T) {
	tests := []struct {
		name             string
		args             Book
		wantedStatements map[string]string
	}{
		{name: "One book",
			args: Book{
				ID:    1,
				Title: "A Hora da Estrela",
				Year:  1977,
				Author: Author{
					ID: 1,
				},
				Pages: 90,
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Livros(titulo, ano, autor, paginas) values (\"A Hora da Estrela\", \"1977\", \"1\", \"90\")",
				"UPDATE": "UPDATE Livros SET titulo=\"A Hora da Estrela\", ano=\"1977\", autor=\"1\", paginas=\"90\" WHERE idLivro = \"1\"",
				"DELETE": "DELETE FROM Livros WHERE idLivro = \"1\"",
				"SELECT": "SELECT * FROM Livros WHERE idLivro = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty title",
			args: Book{
				ID:    1,
				Title: "",
				Year:  1977,
				Author: Author{
					ID: 1,
				},
				Pages: 90,
			},
			wantedStatements: map[string]string{
				"INSERT": "",
				"UPDATE": "",
				"DELETE": "DELETE FROM Livros WHERE idLivro = \"1\"",
				"SELECT": "SELECT * FROM Livros WHERE idLivro = \"1\"",
				"TEST":   "",
			},
		},
		{name: "Empty id",
			args: Book{
				ID:    0,
				Title: "A Hora da Estrela",
				Year:  1977,
				Author: Author{
					ID: 1,
				},
				Pages: 90,
			},
			wantedStatements: map[string]string{
				"INSERT": "INSERT INTO Livros(titulo, ano, autor, paginas) values (\"A Hora da Estrela\", \"1977\", \"1\", \"90\")",
				"UPDATE": "",
				"DELETE": "",
				"SELECT": "",
				"TEST":   "",
			},
		},
		{name: "Empty fields",
			args: Book{
				ID:    0,
				Title: "",
				Year:  0,
				Author: Author{
					ID: 0,
				},
				Pages: 0,
			},
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
