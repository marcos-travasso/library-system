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

	Mock.ExpectExec("INSERT INTO emprestimos").WithArgs(l.Book.ID, l.User.ID, l.LendDay).
		WillReturnResult(sqlmock.NewResult(d.LendingId, 1))
	Mock.ExpectExec("INSERT INTO devolucoes").WithArgs(d.LendingId, l.Devolution.Date).
		WillReturnResult(sqlmock.NewResult(d.DevolutionId, 1))

	err := InsertLending(l)

	require.NoError(t, err)
	require.Equal(t, d.LendingId, l.ID)
	require.Equal(t, d.DevolutionId, l.Devolution.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
