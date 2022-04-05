package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
	"github.com/stretchr/testify/require"
	"testing"
)

type DummyBookParams struct {
	book     models.Book
	bookId   int64
	authorId int64
	genreId  int64
}

func generateValidBook() *DummyBookParams {
	var d DummyBookParams
	b := util.RandomBook()

	d.bookId = b.ID
	d.authorId = b.Author.ID
	d.genreId = b.Genre.ID

	b.ID = 0
	b.Author.ID = 0
	b.Genre.ID = 0
	d.book = b

	return &d
}

func Test_InsertBook_ValidBook(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidBook()
	a := &d.book.Author
	g := &d.book.Genre
	b := &d.book

	//Author queries
	mock.ExpectQuery("SELECT").WithArgs(a.Person.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT INTO Pessoas").WithArgs(a.Person.Name, a.Person.Gender, a.Person.Birthday).
		WillReturnResult(sqlmock.NewResult(a.Person.ID, 1))
	mock.ExpectExec("INSERT INTO Autores").WithArgs(a.Person.ID).
		WillReturnResult(sqlmock.NewResult(d.authorId, 1))

	//Genre queries
	mock.ExpectQuery("SELECT").WithArgs(g.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT INTO Generos").WithArgs(g.Name).
		WillReturnResult(sqlmock.NewResult(d.genreId, 1))

	//Book queries
	mock.ExpectExec("INSERT INTO Livros").WithArgs(b.Title, b.Year, d.authorId, b.Pages).
		WillReturnResult(sqlmock.NewResult(d.bookId, 1))
	mock.ExpectExec("INSERT INTO generos_dos_livros").WithArgs(d.bookId, d.genreId).
		WillReturnResult(sqlmock.NewResult(d.bookId, 1))

	err := InsertBook(b)

	require.NoError(t, err)
	require.Equal(t, d.bookId, b.ID)
	require.Equal(t, d.genreId, g.ID)
	require.Equal(t, d.authorId, a.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
