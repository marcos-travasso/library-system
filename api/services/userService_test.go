package services

import (
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
	//TODO
	return
}

//TODO more tests
