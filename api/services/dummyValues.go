package services

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

type DummyBookParams struct {
	Book         models.Book
	BookId       int64
	AuthorId     int64
	GenreId      int64
	BookRow      *sqlmock.Rows
	LinkGenreRow *sqlmock.Rows
	GenreRow     *sqlmock.Rows
	AuthorRow    *sqlmock.Rows
	PersonRow    *sqlmock.Rows
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

func GenerateValidBook() *DummyBookParams {
	var d DummyBookParams
	b := util.RandomBook()

	d.BookId = b.ID
	d.AuthorId = b.Author.ID
	d.GenreId = b.Genre.ID

	d.BookRow = sqlmock.NewRows([]string{"", "", "", "", ""}).
		AddRow(b.ID, b.Title, b.Year, b.Author.ID, b.Pages)
	g := b.Genre
	d.LinkGenreRow = sqlmock.NewRows([]string{"", ""}).
		AddRow(b.ID, g.ID)
	d.GenreRow = sqlmock.NewRows([]string{"", ""}).
		AddRow(g.ID, g.Name)
	a := b.Author
	d.AuthorRow = sqlmock.NewRows([]string{"", ""}).
		AddRow(a.ID, a.Person.ID)
	p := a.Person
	d.PersonRow = sqlmock.NewRows([]string{"", "", "", ""}).
		AddRow(p.ID, p.Name, p.Gender, p.Birthday)

	b.ID = 0
	b.Author.ID = 0
	b.Genre.ID = 0
	d.Book = b

	return &d
}
