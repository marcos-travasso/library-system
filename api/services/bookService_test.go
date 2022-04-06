package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/fixtures"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InsertBook_ValidBook(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidBook()
	d.MockInsertValues(Mock)

	err := InsertBook(&d.Book)

	require.NoError(t, err)
	require.Equal(t, d.BookId, d.Book.ID)
	require.Equal(t, d.GenreId, d.Book.Genre.ID)
	require.Equal(t, d.AuthorId, d.Book.Author.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_SelectBook_ValidBook(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidBook()
	d.MockSelectValues(Mock)

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

func Test_SelectBooks(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d1 := fixtures.GenerateValidBook()
	d2 := fixtures.GenerateValidBook()

	Mock.ExpectQuery("SELECT idLivro FROM Livros").
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(d1.BookId).
			AddRow(d2.BookId))

	d1.MockSelectValues(Mock)
	d2.MockSelectValues(Mock)

	books, err := SelectBooks()

	require.NoError(t, err)
	require.Equal(t, 2, len(books))

	require.Equal(t, d1.BookId, books[0].ID)
	require.Equal(t, d1.AuthorId, books[0].Author.ID)
	require.Equal(t, d1.GenreId, books[0].Genre.ID)

	require.Equal(t, d2.BookId, books[1].ID)
	require.Equal(t, d2.AuthorId, books[1].Author.ID)
	require.Equal(t, d2.GenreId, books[1].Genre.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
