package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/fixtures"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InsertLending_ValidLending(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidLending()
	l := &d.Lending
	Mock.ExpectQuery("SELECT livro, usuario").WithArgs(l.Book.ID, l.User.ID).
		WillReturnRows(sqlmock.NewRows([]string{"", ""}))
	d.MockInsertValues(Mock)

	err := InsertLending(l)

	require.NoError(t, err)
	require.Equal(t, d.LendingId, l.ID)
	require.Equal(t, d.DevolutionId, l.Devolution.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertLending_BookIsLending(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidLending()
	l := &d.Lending
	Mock.ExpectQuery("SELECT livro, usuario").WithArgs(l.Book.ID, l.User.ID).
		WillReturnRows(sqlmock.NewRows([]string{"", ""}).AddRow(l.Book.ID, -1))

	err := InsertLending(l)

	require.EqualError(t, err, "book already have lending")

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertLending_UserHasLending(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidLending()
	l := &d.Lending
	Mock.ExpectQuery("SELECT livro, usuario").WithArgs(l.Book.ID, l.User.ID).
		WillReturnRows(sqlmock.NewRows([]string{"", ""}).AddRow(-1, l.User.ID))

	err := InsertLending(l)

	require.EqualError(t, err, "user already have lending")

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_ReturnLending_ValidReturn(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidLending()
	l := &d.Lending
	Mock.ExpectExec("UPDATE emprestimos SET").WithArgs(l.ID).
		WillReturnResult(sqlmock.NewResult(l.ID, 0))

	err := ReturnLending(l)

	require.NoError(t, err)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
