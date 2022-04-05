package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InsertBook_ValidBook(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := GenerateValidBook()
	a := &d.Book.Author
	g := &d.Book.Genre
	b := &d.Book

	//Author queries
	Mock.ExpectQuery("SELECT").WithArgs(a.Person.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	Mock.ExpectExec("INSERT INTO Pessoas").WithArgs(a.Person.Name, a.Person.Gender, a.Person.Birthday).
		WillReturnResult(sqlmock.NewResult(a.Person.ID, 1))
	Mock.ExpectExec("INSERT INTO Autores").WithArgs(a.Person.ID).
		WillReturnResult(sqlmock.NewResult(d.AuthorId, 1))

	//Genre queries
	Mock.ExpectQuery("SELECT").WithArgs(g.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	Mock.ExpectExec("INSERT INTO Generos").WithArgs(g.Name).
		WillReturnResult(sqlmock.NewResult(d.GenreId, 1))

	//Book queries
	Mock.ExpectExec("INSERT INTO Livros").WithArgs(b.Title, b.Year, d.AuthorId, b.Pages).
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))
	Mock.ExpectExec("INSERT INTO generos_dos_livros").WithArgs(d.BookId, d.GenreId).
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))

	err := InsertBook(b)

	require.NoError(t, err)
	require.Equal(t, d.BookId, b.ID)
	require.Equal(t, d.GenreId, g.ID)
	require.Equal(t, d.AuthorId, a.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_SelectBook_ValidBook(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := GenerateValidBook()

	Mock.ExpectQuery("SELECT \\* FROM Livros").WithArgs(d.BookId).
		WillReturnRows(d.BookRow)
	Mock.ExpectQuery("SELECT \\* FROM generos_dos_livros").WithArgs(d.BookId).
		WillReturnRows(d.LinkGenreRow)
	Mock.ExpectQuery("SELECT \\* FROM Generos").WithArgs(d.GenreId).
		WillReturnRows(d.GenreRow)
	Mock.ExpectQuery("SELECT \\* FROM Autores").WithArgs(d.AuthorId).
		WillReturnRows(d.AuthorRow)
	Mock.ExpectQuery("SELECT \\* FROM Pessoas").WithArgs(d.Book.Author.Person.ID).
		WillReturnRows(d.PersonRow)

	d.Book.ID = d.BookId
	err := SelectBook(&d.Book)

	require.NoError(t, err)
	require.Equal(t, d.BookId, d.Book.ID)
	require.Equal(t, d.GenreId, d.Book.Genre.ID)
	require.Equal(t, d.AuthorId, d.Book.Author.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
	return
}
