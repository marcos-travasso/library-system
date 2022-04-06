package fixtures

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/util"
)

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

func (d *DummyBookParams) MockInsertValues(mock sqlmock.Sqlmock) {
	//Author queries
	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT INTO Pessoas").
		WillReturnResult(sqlmock.NewResult(d.Book.Author.Person.ID, 1))
	mock.ExpectExec("INSERT INTO Autores").
		WillReturnResult(sqlmock.NewResult(d.AuthorId, 1))

	//Genre queries
	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT INTO Generos").
		WillReturnResult(sqlmock.NewResult(d.GenreId, 1))

	//Book queries
	mock.ExpectExec("INSERT INTO Livros").
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))
	mock.ExpectExec("INSERT INTO generos_dos_livros").
		WillReturnResult(sqlmock.NewResult(d.BookId, 1))
}

func (d *DummyBookParams) MockSelectValues(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("SELECT \\* FROM Livros").
		WillReturnRows(d.BookRow)
	mock.ExpectQuery("SELECT \\* FROM generos_dos_livros").
		WillReturnRows(d.LinkGenreRow)
	mock.ExpectQuery("SELECT \\* FROM Generos").
		WillReturnRows(d.GenreRow)
	mock.ExpectQuery("SELECT \\* FROM Autores").
		WillReturnRows(d.AuthorRow)
	mock.ExpectQuery("SELECT \\* FROM Pessoas").
		WillReturnRows(d.PersonRow)
}
