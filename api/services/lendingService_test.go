package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
	"github.com/stretchr/testify/require"
	"testing"
)

type DummyLendingParams struct {
	lending      models.Lending
	lendingId    int64
	devolutionId int64
}

func generateValidLending() *DummyLendingParams {
	var d DummyLendingParams
	l := util.RandomLending()

	d.lendingId = l.ID
	d.devolutionId = l.Devolution.ID

	l.ID = 0
	l.Devolution.ID = 0
	d.lending = l

	return &d
}

func Test_InsertLending_ValidLending(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidLending()
	l := &d.lending

	mock.ExpectExec("INSERT INTO emprestimos").WithArgs(l.Book.ID, l.User.ID, l.LendDay).
		WillReturnResult(sqlmock.NewResult(d.lendingId, 1))
	mock.ExpectExec("INSERT INTO devolucoes").WithArgs(d.lendingId, l.Devolution.Date).
		WillReturnResult(sqlmock.NewResult(d.devolutionId, 1))

	err := InsertLending(l)

	require.NoError(t, err)
	require.Equal(t, d.lendingId, l.ID)
	require.Equal(t, d.devolutionId, l.Devolution.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
