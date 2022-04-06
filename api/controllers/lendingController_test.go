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
	"testing"
)

func Test_PostLending_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := fixtures.GenerateValidLending()
	l := &d.Lending
	services.Mock.ExpectQuery("SELECT livro, usuario").WithArgs(l.Book.ID, l.User.ID).
		WillReturnRows(sqlmock.NewRows([]string{"", ""}))
	d.MockInsertValues(services.Mock)

	lendingBody, _ := json.Marshal(d.Lending)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/lendings", bytes.NewReader(lendingBody))

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedLending models.Lending
	_ = json.Unmarshal(w.Body.Bytes(), &receivedLending)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, d.LendingId, receivedLending.ID)
	require.Equal(t, d.DevolutionId, receivedLending.Devolution.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

//TODO test haslending (http response must be different)
