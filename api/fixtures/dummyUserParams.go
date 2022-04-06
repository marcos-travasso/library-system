package fixtures

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
)

type DummyUserParams struct {
	User       models.User
	UserId     int64
	PersonId   int64
	AddressId  int64
	UserRow    *sqlmock.Rows
	AddressRow *sqlmock.Rows
	PersonRow  *sqlmock.Rows
}

func GenerateValidUser() *DummyUserParams {
	var d DummyUserParams
	u := util.RandomUser()

	d.UserId = u.ID
	d.PersonId = u.Person.ID
	d.AddressId = u.Address.ID
	d.UserRow = sqlmock.NewRows([]string{"", "", "", "", "", "", "", "", ""}).
		AddRow(u.ID, u.Person.ID, u.CellNumber, u.PhoneNumber, u.Address.ID, u.CPF, u.Email, 0, u.CreationDate)
	d.AddressRow = sqlmock.NewRows([]string{"", "", "", "", "", "", ""}).
		AddRow(u.Address.ID, u.Address.CEP, u.Address.City, u.Address.Neighborhood, u.Address.Street, u.Address.Number, u.Address.Complement)
	d.PersonRow = sqlmock.NewRows([]string{"", "", "", ""}).
		AddRow(u.Person.ID, u.Person.Name, u.Person.Gender, u.Person.Birthday)

	u.ID = 0
	u.Person.ID = 0
	u.Address.ID = 0
	d.User = u

	return &d
}

func (d *DummyUserParams) MockInsertValues(mock sqlmock.Sqlmock) {
	p := &d.User.Person
	a := &d.User.Address
	u := &d.User

	mock.ExpectExec("INSERT INTO Pessoas").WithArgs(p.Name, p.Gender, p.Birthday).
		WillReturnResult(sqlmock.NewResult(d.PersonId, 1))
	mock.ExpectExec("INSERT INTO Enderecos").WithArgs(a.CEP, a.City, a.Neighborhood, a.Street, a.Number, a.Complement).
		WillReturnResult(sqlmock.NewResult(d.AddressId, 1))
	mock.ExpectExec("INSERT INTO Usuarios").WithArgs(d.PersonId, u.CellNumber, u.PhoneNumber, d.AddressId, u.CPF, u.Email).
		WillReturnResult(sqlmock.NewResult(d.UserId, 1))
}

func (d *DummyUserParams) MockSelectValues(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("SELECT \\* FROM Usuarios").WithArgs(d.UserId).
		WillReturnRows(d.UserRow)
	mock.ExpectQuery("SELECT \\* FROM Enderecos").WithArgs(d.AddressId).
		WillReturnRows(d.AddressRow)
	mock.ExpectQuery("SELECT \\* FROM Pessoas").WithArgs(d.PersonId).
		WillReturnRows(d.PersonRow)
}
