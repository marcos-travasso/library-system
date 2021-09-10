package database

import (
	"fmt"
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
	"time"
)

var booksTests = []struct {
	name string
	args structs.Book
}{
	{
		name: "First book",
		args: structs.Book{
			ID:    1,
			Year:  1977,
			Pages: 90,
			Title: "A hora da estrela",
			Author: structs.Author{
				ID: 1,
				Person: structs.Person{
					ID:       1,
					Name:     "Clarice Lispector",
					Gender:   "F",
					Birthday: "10/12/1920",
				},
			},
			Genre: structs.Genre{
				ID:   1,
				Name: "Romance",
			},
		},
	},
	{
		name: "Second book",
		args: structs.Book{
			ID:    2,
			Year:  1881,
			Pages: 368,
			Title: "Memórias, Póstumas de Brás Cubas",
			Author: structs.Author{
				ID: 2,
				Person: structs.Person{
					ID:       2,
					Name:     "Machado de Assis",
					Gender:   "M",
					Birthday: "21/6/1839",
				},
			},
			Genre: structs.Genre{
				ID:   1,
				Name: "Romance",
			},
		},
	},
}

var usersTests = []struct {
	name string
	args structs.User
}{
	{
		name: "First user",
		args: structs.User{
			ID: 1,
			Person: structs.Person{
				ID:       1,
				Name:     "Marcos",
				Birthday: "01/01/2002",
				Gender:   "M",
			},
			CellNumber:  "12345678910",
			PhoneNumber: "9876543210",
			CPF:         "12345678910",
			Email:       "testmail@mail.test",
			Address: structs.Address{
				ID:           1,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Boqueirão",
				Street:       "XV",
				Complement:   "",
			},
		},
	},
	{
		name: "Second user",
		args: structs.User{
			ID: 2,
			Person: structs.Person{
				ID:       2,
				Name:     "Fulano",
				Birthday: "10/10/2000",
				Gender:   "M",
			},
			CellNumber:  "64815973281",
			PhoneNumber: "8569459461",
			CPF:         "31648519624",
			Email:       "fulano@mail.test",
			Address: structs.Address{
				ID:           2,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Centro",
				Street:       "XV",
				Complement:   "",
			},
		},
	},
	{
		name: "Third user",
		args: structs.User{
			ID: 3,
			Person: structs.Person{
				ID:       3,
				Name:     "Ciclana",
				Birthday: "05/12/1999",
				Gender:   "F",
			},
			CellNumber:  "64815379281",
			PhoneNumber: "8569459461",
			CPF:         "31648519624",
			Email:       "ciclana@mail.test",
			Address: structs.Address{
				ID:           3,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Centro",
				Street:       "XV",
				Complement:   "",
			},
		},
	},
}

func TestDatabase_InsertLending(t *testing.T) {
	currentTime := time.Now()

	dbDir := Database{Dir: "./temp/test_lendingInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range usersTests {
		tt.args.ID, err = dbDir.InsertUser(tt.args)
		if err != nil {
			t.Error(err)
		}
	}

	for _, tt := range booksTests {
		tt.args.ID, err = dbDir.InsertBook(tt.args)
		if err != nil {
			t.Error(err)
		}
	}

	lendings := make([]structs.Lending, len(usersTests)+3, len(usersTests)+3)

	for i := 0; i < 2; i++ {
		t.Run(fmt.Sprintf("Lending %d", i), func(t *testing.T) {
			lendings[i] = structs.Lending{
				User:    usersTests[i].args,
				Book:    booksTests[i].args,
				LendDay: currentTime.Format("2006-01-02"),
				Devolution: []structs.Devolution{
					{
						ID:   i + 1,
						Date: "31/10/2021",
					},
				},
			}

			lendings[i].ID, err = dbDir.InsertLending(lendings[i])

			if err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("User already has lending", func(t *testing.T) {
		lendings[2] = structs.Lending{
			User:    usersTests[0].args,
			Book:    booksTests[0].args,
			LendDay: currentTime.Format("2006-01-02"),
			Devolution: []structs.Devolution{
				{
					ID:   3,
					Date: "31/10/2021",
				},
			},
		}

		lendings[2].ID, err = dbDir.InsertLending(lendings[2])

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Book already has lending", func(t *testing.T) {
		lendings[3] = structs.Lending{
			User:    usersTests[2].args,
			Book:    booksTests[0].args,
			LendDay: currentTime.Format("2006-01-02"),
			Devolution: []structs.Devolution{
				{
					ID:   3,
					Date: "31/10/2021",
				},
			},
		}

		lendings[3].ID, err = dbDir.InsertLending(lendings[3])

		if err == nil {
			t.Error(err)
		}
	})
}
