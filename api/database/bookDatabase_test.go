package database

import (
	"encoding/json"
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
)

func TestDatabase_InsertBook(t *testing.T) {
	tests := []struct {
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
		{
			name: "Third book",
			args: structs.Book{
				ID:    3,
				Year:  1882,
				Pages: 96,
				Title: "O Alienista",
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
					ID:   2,
					Name: "Tale",
				},
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_bookInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := dbDir.InsertBook(tt.args)
			if err != nil {
				t.Error(err)
			}

			if id != tt.args.ID {
				t.Errorf("Fail to check book id, got %d expect %d", id, tt.args.ID)
			}
		})
	}
}

func TestDatabase_SelectBook(t *testing.T) {
	tests := []struct {
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
		{
			name: "Third book",
			args: structs.Book{
				ID:    3,
				Year:  1882,
				Pages: 96,
				Title: "O Alienista",
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
					ID:   2,
					Name: "Tale",
				},
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_bookSelect.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dbDir.InsertBook(tt.args)
			if err != nil {
				t.Error(err)
			}

			got, err := dbDir.SelectBook(tt.args)
			if got != tt.args {
				gotJSON, _ := json.Marshal(got)
				wantJSON, _ := json.Marshal(tt.args)

				t.Errorf("SelectBook() got = %v, want %v", string(gotJSON), string(wantJSON))
			}
		})
	}
}

func TestDatabase_SelectBooks(t *testing.T) {
	{
		tests := []struct {
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
			{
				name: "Third book",
				args: structs.Book{
					ID:    3,
					Year:  1882,
					Pages: 96,
					Title: "O Alienista",
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
						ID:   2,
						Name: "Tale",
					},
				},
			},
		}

		dbDir := Database{Dir: "./temp/test_booksSelect.db"}
		err := dbDir.clearDatabase()
		if err != nil {
			t.Error(err)
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := dbDir.InsertBook(tt.args)
				if err != nil {
					t.Error(err)
				}
			})
		}

		books := make([]structs.Book, len(tests), len(tests))
		books, err = dbDir.SelectBooks()
		if err != nil {
			t.Error(err)
		}

		for i := range tests {
			if tests[i].args != books[i] {
				gotJSON, _ := json.Marshal(books[i])
				wantJSON, _ := json.Marshal(tests[i].args)

				t.Errorf("Select books got = %v, want %v", string(gotJSON), string(wantJSON))
			}
		}
	}
}
