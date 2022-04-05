package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
	"github.com/stretchr/testify/require"
	"testing"
)

type DummyUserParams struct {
	user       models.User
	userId     int64
	personId   int64
	addressId  int64
	userRow    *sqlmock.Rows
	addressRow *sqlmock.Rows
	personRow  *sqlmock.Rows
}

func generateValidUser() *DummyUserParams {
	var d DummyUserParams
	u := util.RandomUser()

	d.userId = u.ID
	d.personId = u.Person.ID
	d.addressId = u.Address.ID
	d.userRow = sqlmock.NewRows([]string{"", "", "", "", "", "", "", "", ""}).
		AddRow(u.ID, u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email, 0, u.CreationDate)
	d.addressRow = sqlmock.NewRows([]string{"", "", "", "", "", "", ""}).
		AddRow(u.Address.ID, u.Address.CEP, u.Address.City, u.Address.Neighborhood, u.Address.Street, u.Address.Number, u.Address.Complement)
	d.personRow = sqlmock.NewRows([]string{"", "", "", ""}).
		AddRow(u.Person.ID, u.Person.Name, u.Person.Gender, u.Person.Birthday)

	u.ID = 0
	u.Person.ID = 0
	u.Address.ID = 0
	d.user = u

	return &d
}

func Test_InsertUser_ValidUser(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidUser()
	p := &d.user.Person
	a := &d.user.Address
	u := &d.user

	mock.ExpectExec("INSERT INTO Pessoas").WithArgs(p.Name, p.Gender, p.Birthday).
		WillReturnResult(sqlmock.NewResult(d.personId, 1))
	mock.ExpectExec("INSERT INTO Enderecos").WithArgs(a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement).
		WillReturnResult(sqlmock.NewResult(d.addressId, 1))
	mock.ExpectExec("INSERT INTO Usuarios").WithArgs(d.personId, u.CellNumber, u.PhoneNumber, d.addressId, u.CPF, u.Email).
		WillReturnResult(sqlmock.NewResult(d.userId, 1))

	err := InsertUser(u)

	require.NoError(t, err)
	require.Equal(t, d.personId, u.Person.ID)
	require.Equal(t, d.addressId, u.Address.ID)
	require.Equal(t, d.userId, u.ID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_InsertUser_InvalidUser(t *testing.T) {
	//TODO
	return
}

func Test_SelectUser_ValidUser(t *testing.T) {
	InitializeTestServices()
	defer db.Close()

	d := generateValidUser()

	mock.ExpectQuery("SELECT \\* FROM Usuarios").WithArgs(d.userId).
		WillReturnRows(d.userRow)
	mock.ExpectQuery("SELECT \\* FROM Enderecos").WithArgs(d.addressId).
		WillReturnRows(d.addressRow)
	mock.ExpectQuery("SELECT \\* FROM Pessoas").WithArgs(d.personId).
		WillReturnRows(d.personRow)

	d.user.ID = d.userId
	err := SelectUser(&d.user)

	require.NoError(t, err)
	require.Equal(t, d.userId, d.user.ID)
	require.Equal(t, d.personId, d.user.Person.ID)
	require.Equal(t, d.addressId, d.user.Address.ID)

	err = mock.ExpectationsWereMet()
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
