package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/fixtures"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InsertUser_ValidUser(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidUser()
	d.MockInsertValues(Mock)

	err := InsertUser(&d.User)

	require.NoError(t, err)
	require.Equal(t, d.PersonId, d.User.Person.ID)
	require.Equal(t, d.AddressId, d.User.Address.ID)
	require.Equal(t, d.UserId, d.User.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertUser_InvalidUser(t *testing.T) {
	//TODO
	return
}

func Test_SelectUser_ValidUser(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := fixtures.GenerateValidUser()
	d.MockSelectValues(Mock)

	d.User.ID = d.UserId
	err := SelectUser(&d.User)

	require.NoError(t, err)
	require.Equal(t, d.UserId, d.User.ID)
	require.Equal(t, d.PersonId, d.User.Person.ID)
	require.Equal(t, d.AddressId, d.User.Address.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_SelectUser_InvalidUser(t *testing.T) {
	//TODO
	return
}

func Test_SelectUsers(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d1 := fixtures.GenerateValidUser()
	d2 := fixtures.GenerateValidUser()

	Mock.ExpectQuery("SELECT idUsuario FROM Usuarios").
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(d1.UserId).
			AddRow(d2.UserId))

	d1.MockSelectValues(Mock)
	d2.MockSelectValues(Mock)

	users, err := SelectUsers()

	require.NoError(t, err)
	require.Equal(t, 2, len(users))

	require.Equal(t, d1.UserId, users[0].ID)
	require.Equal(t, d1.PersonId, users[0].Person.ID)
	require.Equal(t, d1.AddressId, users[0].Address.ID)

	require.Equal(t, d2.UserId, users[1].ID)
	require.Equal(t, d2.PersonId, users[1].Person.ID)
	require.Equal(t, d2.AddressId, users[1].Address.ID)

	err = Mock.ExpectationsWereMet()
	require.NoError(t, err)
}

//TODO more tests
