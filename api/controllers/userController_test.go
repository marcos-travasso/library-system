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

func Test_PostUser_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := services.GenerateValidUser()
	userBody, _ := json.Marshal(d.User)

	services.Mock.ExpectExec("INSERT INTO Pessoas").
		WillReturnResult(sqlmock.NewResult(d.PersonId, 1))
	services.Mock.ExpectExec("INSERT INTO Enderecos").
		WillReturnResult(sqlmock.NewResult(d.AddressId, 1))
	services.Mock.ExpectExec("INSERT INTO Usuarios").
		WillReturnResult(sqlmock.NewResult(d.UserId, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(userBody))

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedUser models.User
	_ = json.Unmarshal(w.Body.Bytes(), &receivedUser)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, d.UserId, receivedUser.ID)
	require.Equal(t, d.AddressId, receivedUser.Address.ID)
	require.Equal(t, d.PersonId, receivedUser.Person.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
