package structs

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int    `json:"id"`
	Year   int    `json:"year"`
	Pages  int    `json:"pages"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (g Genre) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if g.Name == "" {
			return "", errors.New("genre has no name")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Generos(nome) values (\"%s\")", g.Name)
	case "UPDATE":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		if g.Name == "" {
			return "", errors.New("genre has no name")
		}
		sqlStatement += fmt.Sprintf("UPDATE Generos SET nome=\"%s\" WHERE idGenero = \"%d\"", g.Name, g.ID)
	case "DELETE":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Generos WHERE idGenero = \"%d\"", g.ID)
	case "SELECT":
		if g.ID == 0 {
			return "", errors.New("genre has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Generos WHERE idGenero = \"%d\"", g.ID)
	case "EXIST":
		if g.Name == "" {
			return "", errors.New("genre has no name")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Generos WHERE nome = \"%s\"", g.Name)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}

func (b Book) SQLStatement(statementType string) (string, error) {
	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if b.Title == "" {
			return "", errors.New("book has no title")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO Livros(titulo, ano, autor, paginas) values (\"%s\", \"%d\", \"%d\", \"%d\")", b.Title, b.Year, b.Author.ID, b.Pages)
	case "UPDATE":
		if b.ID == 0 {
			return "", errors.New("book has no id")
		}
		if b.Title == "" {
			return "", errors.New("book has no title")
		}
		sqlStatement += fmt.Sprintf("UPDATE Livros SET titulo=\"%s\", ano=\"%d\", autor=\"%d\", paginas=\"%d\" WHERE idLivro = \"%d\"", b.Title, b.Year, b.Author.ID, b.Pages, b.ID)
	case "DELETE":
		if b.ID == 0 {
			return "", errors.New("book has no ID")
		}
		sqlStatement += fmt.Sprintf("DELETE FROM Livros WHERE idLivro = \"%d\"", b.ID)
	case "SELECT":
		if b.ID == 0 {
			return "", errors.New("book has no ID")
		}
		sqlStatement += fmt.Sprintf("SELECT * FROM Livros WHERE idLivro = \"%d\"", b.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}

func (b Book) LinkSQLStatement(statementType string) (string, error) {
	if b.ID == 0 {
		return "", errors.New("book has no id")
	}

	sqlStatement := ""

	switch statementType {
	case "INSERT":
		if b.Genre.ID == 0 {
			return "", errors.New("genre has no id")
		}
		sqlStatement += fmt.Sprintf("INSERT INTO generos_dos_livros(livro, genero) values (\"%d\", \"%d\")", b.ID, b.Genre.ID)
	case "UPDATE":
		if b.Genre.ID == 0 {
			return "", errors.New("genre has no id")
		}
		sqlStatement += fmt.Sprintf("UPDATE generos_dos_livros SET genero=\"%d\" WHERE livro = \"%d\"", b.Genre.ID, b.ID)
	case "DELETE":
		sqlStatement += fmt.Sprintf("DELETE FROM generos_dos_livros WHERE livro = \"%d\"", b.ID)
	case "SELECT":
		sqlStatement += fmt.Sprintf("SELECT * FROM generos_dos_livros WHERE livro = \"%d\"", b.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}

func (g Genre) LinkSQLStatement(statementType string) (string, error) {
	if g.ID == 0 {
		return "", errors.New("genre has no id")
	}

	sqlStatement := ""

	switch statementType {
	case "DELETE":
		sqlStatement += fmt.Sprintf("DELETE FROM generos_dos_livros WHERE genero = \"%d\"", g.ID)
	case "SELECT":
		sqlStatement += fmt.Sprintf("SELECT * FROM generos_dos_livros WHERE genero = \"%d\"", g.ID)
	default:
		return "", errors.New("invalid statement type")
	}
	return sqlStatement, nil
}
