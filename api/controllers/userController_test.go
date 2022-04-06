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

func Test_PostUser_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := fixtures.GenerateValidUser()
	d.MockInsertValues(services.Mock)

	userBody, _ := json.Marshal(d.User)

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

func Test_GetUser_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d := fixtures.GenerateValidUser()
	d.MockSelectValues(services.Mock)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(int(d.UserId)), nil)

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

func Test_GetUsers_ShouldReturnOk(t *testing.T) {
	//GIVEN
	InitializeControllers()
	services.InitializeTestServices()

	d1 := fixtures.GenerateValidUser()
	d2 := fixtures.GenerateValidUser()

	services.Mock.ExpectQuery("SELECT idUsuario FROM Usuarios").
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(d1.UserId).
			AddRow(d2.UserId))

	d1.MockSelectValues(services.Mock)
	d2.MockSelectValues(services.Mock)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)

	//WHEN
	router.ServeHTTP(w, req)

	//THEN
	var receivedUsers []models.User
	_ = json.Unmarshal(w.Body.Bytes(), &receivedUsers)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, 2, len(receivedUsers))

	require.Equal(t, d1.UserId, receivedUsers[0].ID)
	require.Equal(t, d1.PersonId, receivedUsers[0].Person.ID)
	require.Equal(t, d1.AddressId, receivedUsers[0].Address.ID)

	require.Equal(t, d2.UserId, receivedUsers[1].ID)
	require.Equal(t, d2.PersonId, receivedUsers[1].Person.ID)
	require.Equal(t, d2.AddressId, receivedUsers[1].Address.ID)

	err := services.Mock.ExpectationsWereMet()
	require.NoError(t, err)
}
