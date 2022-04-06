package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/services"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_PostBook_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := services.GenerateValidBook()
	bookBody, _ := json.Marshal(d.Book)

	//Author queries
	services.Mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{}))
	services.Mock.ExpectExec("INSERT INTO Pessoas").
		WillReturnResult(sqlmock.NewResult(d.Book.Author.Person.ID, 1))
	services.Mock.ExpectExec("INSERT INTO Autores").
		WillReturnResult(sqlmock.NewResult(d.AuthorId, 1))

	//Genre queries
	services.Mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{}))
	services.Mock.ExpectExec("INSERT INTO Generos").
		WillReturnResult(sqlmock.NewResult(d.GenreId, 1))

	//Book queries
	services.Mock.ExpectExec("INSERT INTO Livros").
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))
	services.Mock.ExpectExec("INSERT INTO generos_dos_livros").
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(bookBody))

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedBook models.Book
	_ = json.Unmarshal(w.Body.Bytes(), &receivedBook)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, d.BookId, receivedBook.ID)
	require.Equal(t, d.AuthorId, receivedBook.Author.ID)
	require.Equal(t, d.GenreId, receivedBook.Genre.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_GetBook_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := services.GenerateValidBook()

	services.Mock.ExpectQuery("SELECT \\* FROM Livros").
		WillReturnRows(d.BookRow)
	services.Mock.ExpectQuery("SELECT \\* FROM generos_dos_livros").
		WillReturnRows(d.LinkGenreRow)
	services.Mock.ExpectQuery("SELECT \\* FROM Generos").
		WillReturnRows(d.GenreRow)
	services.Mock.ExpectQuery("SELECT \\* FROM Autores").
		WillReturnRows(d.AuthorRow)
	services.Mock.ExpectQuery("SELECT \\* FROM Pessoas").
		WillReturnRows(d.PersonRow)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(int(d.BookId)), nil)

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedBook models.Book
	_ = json.Unmarshal(w.Body.Bytes(), &receivedBook)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, d.BookId, receivedBook.ID)
	require.Equal(t, d.GenreId, receivedBook.Genre.ID)
	require.Equal(t, d.AuthorId, receivedBook.Author.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
