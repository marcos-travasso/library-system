package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
	"github.com/stretchr/testify/require"
	"testing"
)

type DummyAuthorParams struct {
	author    models.Author
	authorId  int64
	personId  int64
	authorRow *sqlmock.Rows
	personRow *sqlmock.Rows
}

func generateValidAuthor() *DummyAuthorParams {
	var d DummyAuthorParams
	a := util.RandomAuthor()

	d.authorId = a.ID
	d.personId = a.Person.ID
	p := a.Person
	d.authorRow = sqlmock.NewRows([]string{"", ""}).AddRow(a.ID, a.Person.ID)
	d.personRow = sqlmock.NewRows([]string{"", "", "", ""}).AddRow(p.ID, p.Name, p.Gender, p.Birthday)

	a.ID = 0
	a.Person.ID = 0
	d.author = a

	return &d
}

func Test_InsertAuthor_ValidAuthor(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidAuthor()
	p := &d.author.Person
	a := &d.author

	Mock.ExpectQuery("SELECT").WithArgs(p.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	Mock.ExpectExec("INSERT INTO Pessoas").WithArgs(p.Name, p.Gender, p.Birthday).
		WillReturnResult(sqlmock.NewResult(d.personId, 1))
	Mock.ExpectExec("INSERT INTO Autores").WithArgs(d.personId).
		WillReturnResult(sqlmock.NewResult(d.authorId, 1))

	err := InsertAuthor(a)

	require.NoError(t, err)
	require.Equal(t, d.personId, p.ID)
	require.Equal(t, d.authorId, a.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertAuthor_AlreadyInserted(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidAuthor()
	p := &d.author.Person
	a := &d.author

	Mock.ExpectQuery("SELECT").WithArgs(p.Name).
		WillReturnRows(sqlmock.NewRows([]string{"idAutor", "pessoa"}).AddRow(d.authorId, d.personId))

	err := InsertAuthor(a)

	require.NoError(t, err)
	require.Equal(t, d.personId, a.Person.ID)
	require.Equal(t, d.authorId, a.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_SelectAuthor_ValidAuthor(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidAuthor()
	p := &d.author.Person
	a := &d.author

	Mock.ExpectQuery("SELECT \\* FROM Autores").
		WillReturnRows(d.authorRow)
	Mock.ExpectQuery("SELECT \\* FROM Pessoas").
		WillReturnRows(d.personRow)

	err := SelectAuthor(a)

	require.NoError(t, err)
	require.Equal(t, d.authorId, a.ID)
	require.Equal(t, d.personId, p.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
