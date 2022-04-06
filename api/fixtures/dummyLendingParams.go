package fixtures

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
)

type DummyLendingParams struct {
	Lending      models.Lending
	LendingId    int64
	DevolutionId int64
	BookId       int64
	UserId       int64
}

func GenerateValidLending() *DummyLendingParams {
	var d DummyLendingParams
	l := util.RandomLending()

	d.LendingId = l.ID
	d.DevolutionId = l.Devolution.ID
	d.BookId = l.Book.ID
	d.UserId = l.User.ID

	l.ID = 0
	l.Devolution.ID = 0
	d.Lending = l

	return &d
}

func (d *DummyLendingParams) MockInsertValues(mock sqlmock.Sqlmock) {
	mock.ExpectExec("INSERT INTO emprestimos").
		WillReturnResult(sqlmock.NewResult(d.LendingId, 1))
	mock.ExpectExec("INSERT INTO devolucoes").
		WillReturnResult(sqlmock.NewResult(d.DevolutionId, 1))
}
