package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
	"github.com/stretchr/testify/require"
	"testing"
)

type DummyGenreParams struct {
	genre    models.Genre
	genreId  int64
	genreRow *sqlmock.Rows
}

func generateValidGenre() *DummyGenreParams {
	var d DummyGenreParams
	g := util.RandomGenre()

	d.genreId = g.ID
	d.genreRow = sqlmock.NewRows([]string{"", ""}).AddRow(g.ID, g.Name)

	g.ID = 0
	d.genre = g

	return &d
}

func Test_InsertGenre_ValidGenre(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidGenre()
	g := &d.genre

	mock.ExpectQuery("SELECT").WithArgs(g.Name).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT INTO Generos").WithArgs(g.Name).
		WillReturnResult(sqlmock.NewResult(d.genreId, 1))

	err := InsertGenre(g)

	require.NoError(t, err)
	require.Equal(t, d.genreId, g.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertGenre_AlreadyInserted(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidGenre()
	g := &d.genre

	mock.ExpectQuery("SELECT").WithArgs(g.Name).
		WillReturnRows(sqlmock.NewRows([]string{"idGenero"}).AddRow(d.genreId))

	err := InsertGenre(g)

	require.NoError(t, err)
	require.Equal(t, d.genreId, g.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_SelectGenre_ValidGenre(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidGenre()
	g := &d.genre

	mock.ExpectQuery("SELECT \\* from Generos").
		WillReturnRows(d.genreRow)

	err := SelectGenre(g)

	require.NoError(t, err)
	require.Equal(t, d.genreId, g.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
