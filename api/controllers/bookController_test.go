package controllers

import (
	"bytes"
	"encoding/json"
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
