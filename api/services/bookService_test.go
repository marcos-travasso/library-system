package services

import (
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
