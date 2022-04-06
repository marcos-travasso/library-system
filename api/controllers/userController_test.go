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
