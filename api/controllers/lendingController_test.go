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
	"testing"
)

func Test_PostLending_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := services.GenerateValidLending()
	lendingBody, _ := json.Marshal(d.Lending)

	services.Mock.ExpectExec("INSERT INTO emprestimos").
		WillReturnResult(sqlmock.NewResult(d.LendingId, 1))
	services.Mock.ExpectExec("INSERT INTO devolucoes").
		WillReturnResult(sqlmock.NewResult(d.DevolutionId, 1))

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
