package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InsertUser_ValidUser(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := GenerateValidUser()
	p := &d.User.Person
	a := &d.User.Address
	u := &d.User

	Mock.ExpectExec("INSERT INTO Pessoas").WithArgs(p.Name, p.Gender, p.Birthday).
		WillReturnResult(sqlmock.NewResult(d.PersonId, 1))
	Mock.ExpectExec("INSERT INTO Enderecos").WithArgs(a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement).
		WillReturnResult(sqlmock.NewResult(d.AddressId, 1))
	Mock.ExpectExec("INSERT INTO Usuarios").WithArgs(d.PersonId, u.CellNumber, u.PhoneNumber, d.AddressId, u.CPF, u.Email).
		WillReturnResult(sqlmock.NewResult(d.UserId, 1))

	err := InsertUser(u)

	require.NoError(t, err)
	require.Equal(t, d.PersonId, u.Person.ID)
	require.Equal(t, d.AddressId, u.Address.ID)
	require.Equal(t, d.UserId, u.ID)

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

	d := GenerateValidUser()

	Mock.ExpectQuery("SELECT \\* FROM Usuarios").WithArgs(d.UserId).
		WillReturnRows(d.UserRow)
	Mock.ExpectQuery("SELECT \\* FROM Enderecos").WithArgs(d.AddressId).
		WillReturnRows(d.AddressRow)
	Mock.ExpectQuery("SELECT \\* FROM Pessoas").WithArgs(d.PersonId).
		WillReturnRows(d.PersonRow)

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
