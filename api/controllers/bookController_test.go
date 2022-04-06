package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/fixtures"
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

	d := fixtures.GenerateValidBook()
	d.MockInsertValues(services.Mock)

	bookBody, _ := json.Marshal(d.Book)

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

	d := fixtures.GenerateValidBook()
	d.MockSelectValues(services.Mock)

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

func Test_GetBooks_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d1 := fixtures.GenerateValidBook()
	d2 := fixtures.GenerateValidBook()

	services.Mock.ExpectQuery("SELECT idLivro FROM Livros").
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(d1.BookId).
			AddRow(d2.BookId))

	d1.MockSelectValues(services.Mock)
	d2.MockSelectValues(services.Mock)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedBooks []models.Book
	_ = json.Unmarshal(w.Body.Bytes(), &receivedBooks)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, 2, len(receivedBooks))

	require.Equal(t, d1.BookId, receivedBooks[0].ID)
	require.Equal(t, d1.AuthorId, receivedBooks[0].Author.ID)
	require.Equal(t, d1.GenreId, receivedBooks[0].Genre.ID)

	require.Equal(t, d2.BookId, receivedBooks[1].ID)
	require.Equal(t, d2.AuthorId, receivedBooks[1].Author.ID)
	require.Equal(t, d2.GenreId, receivedBooks[1].Genre.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
