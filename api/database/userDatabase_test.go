package database

import (
	"encoding/json"
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
	"time"
)

func TestDatabase_InsertUser(t *testing.T) {
	tests := []struct {
		name string
		args structs.User
		want int
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
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   "",
				},
			},
			want: 1,
		},
	}

	dbDir := Database{Dir: "./temp/test_userInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbDir.InsertUser(tt.args)
			if err != nil {
				t.Errorf("InsertUser() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("InsertUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_SelectUser(t *testing.T) {
	currentTime := time.Now()

	tests := []struct {
		name    string
		args    structs.User
		wantErr bool
	}{
		{
			name: "First user",
			args: structs.User{
				ID: 0,
				Person: structs.Person{
					ID:       0,
					Name:     "Marcos",
					Birthday: "01/01/2002",
					Gender:   "M",
				},
				CellNumber:  "12345678910",
				PhoneNumber: "9876543210",
				CPF:         "12345678910",
				Email:       "testmail@mail.test",
				Address: structs.Address{
					ID:           0,
					Number:       123,
					CEP:          "12345678",
					City:         "Curitiba",
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   ""},
				CreationDate: currentTime.Format("2006-01-02"),
			},
		},
		{
			name: "No name",
			args: structs.User{
				ID: 0,
				Person: structs.Person{
					ID:       0,
					Name:     "",
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
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   ""},
				CreationDate: currentTime.Format("2006-01-02"),
			},
			wantErr: true,
		},
		{name: "Empty fields",
			args: structs.User{
				ID: 0,
				Person: structs.Person{
					ID:       0,
					Name:     "",
					Birthday: "",
					Gender:   "",
				},
				CellNumber:  "",
				PhoneNumber: "",
				CPF:         "",
				Email:       "",
				Address: structs.Address{
					ID:           0,
					Number:       0,
					CEP:          "",
					City:         "",
					Neighborhood: "",
					Street:       "",
					Complement:   ""},
				CreationDate: "",
			},
			wantErr: true,
		},
		{
			name: "Second user",
			args: structs.User{
				ID: 0,
				Person: structs.Person{
					ID:       0,
					Name:     "Clarice",
					Birthday: "01/01/2000",
					Gender:   "F",
				},
				CellNumber:  "10987654321",
				PhoneNumber: "0123456789",
				CPF:         "10987654321",
				Email:       "testmail@test.mail",
				Address: structs.Address{
					ID:           0,
					Number:       321,
					CEP:          "98765432",
					City:         "Curitiba",
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   ""},
				CreationDate: currentTime.Format("2006-01-02"),
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_userSelect.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}
	createdUserIDs := 0

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dbDir.InsertUser(tt.args)
			if tt.wantErr && err != nil {
				return
			} else if tt.wantErr && err == nil {
				t.Error(err)
			}
			createdUserIDs++
			tt.args.ID, tt.args.Person.ID, tt.args.Address.ID = createdUserIDs, createdUserIDs, createdUserIDs

			got, err := dbDir.SelectUser(tt.args)
			if got != tt.args {
				gotJSON, _ := json.Marshal(got)
				wantJSON, _ := json.Marshal(tt.args)

				t.Errorf("SelectUser() got = %v, want %v", string(gotJSON), string(wantJSON))
			}
		})
	}
}

func TestDatabase_SelectUsers(t *testing.T) {
	{
		currentTime := time.Now()

		tests := []struct {
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
						Neighborhood: "Centro",
						Street:       "XV",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
			{
				name: "Second user",
				args: structs.User{
					ID: 2,
					Person: structs.Person{
						ID:       2,
						Name:     "Clarice",
						Birthday: "01/01/2000",
						Gender:   "F",
					},
					CellNumber:  "10987654321",
					PhoneNumber: "0123456789",
					CPF:         "10987654321",
					Email:       "testmail@test.mail",
					Address: structs.Address{
						ID:           2,
						Number:       321,
						CEP:          "98765432",
						City:         "Curitiba",
						Neighborhood: "Centro",
						Street:       "XV",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
			{
				name: "Third user",
				args: structs.User{
					ID: 3,
					Person: structs.Person{
						ID:       3,
						Name:     "Aldous Huxley",
						Birthday: "01/01/1970",
						Gender:   "M",
					},
					CellNumber:  "",
					PhoneNumber: "",
					CPF:         "",
					Email:       "",
					Address: structs.Address{
						ID:           3,
						Number:       0,
						CEP:          "123456789",
						City:         "",
						Neighborhood: "",
						Street:       "",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
		}

		dbDir := Database{Dir: "./temp/test_usersSelect.db"}
		err := dbDir.clearDatabase()
		if err != nil {
			t.Error(err)
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := dbDir.InsertUser(tt.args)
				if err != nil {
					t.Error(err)
				}
			})
		}

		users := make([]structs.User, len(tests), len(tests))
		users, err = dbDir.SelectUsers()
		if err != nil {
			t.Error(err)
		}

		for i := range tests {
			if tests[i].args != users[i] {
				gotJSON, _ := json.Marshal(users[i])
				wantJSON, _ := json.Marshal(tests[i].args)

				t.Errorf("Select users got = %v, want %v", string(gotJSON), string(wantJSON))
			}
		}
	}
}

func TestDatabase_DeleteUser(t *testing.T) {
	{
		currentTime := time.Now()

		tests := []struct {
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
						Neighborhood: "Centro",
						Street:       "XV",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
			{
				name: "Second user",
				args: structs.User{
					ID: 2,
					Person: structs.Person{
						ID:       2,
						Name:     "Clarice",
						Birthday: "01/01/2000",
						Gender:   "F",
					},
					CellNumber:  "10987654321",
					PhoneNumber: "0123456789",
					CPF:         "10987654321",
					Email:       "testmail@test.mail",
					Address: structs.Address{
						ID:           2,
						Number:       321,
						CEP:          "98765432",
						City:         "Curitiba",
						Neighborhood: "Centro",
						Street:       "XV",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
			{
				name: "Third user",
				args: structs.User{
					ID: 3,
					Person: structs.Person{
						ID:       3,
						Name:     "Aldous Huxley",
						Birthday: "01/01/1970",
						Gender:   "M",
					},
					CellNumber:  "",
					PhoneNumber: "",
					CPF:         "",
					Email:       "",
					Address: structs.Address{
						ID:           3,
						Number:       0,
						CEP:          "123456789",
						City:         "",
						Neighborhood: "",
						Street:       "",
						Complement:   ""},
					CreationDate: currentTime.Format("2006-01-02"),
				},
			},
		}

		dbDir := Database{Dir: "./temp/test_usersDelete.db"}
		err := dbDir.clearDatabase()
		if err != nil {
			t.Error(err)
		}

		for _, tt := range tests {
			_, err := dbDir.InsertUser(tt.args)
			if err != nil {
				t.Error(err)
			}
		}

		usersCountBefore, err := dbDir.countRows("Usuarios")
		if err != nil {
			t.Error(err)
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := dbDir.DeleteUser(tt.args)
				if err != nil {
					t.Error(err)
				}
				usersCountBefore--
			})
		}

		columnsToCount := []string{"Usuarios", "Enderecos", "Pessoas"}
		for _, column := range columnsToCount {
			countAfter, err := dbDir.countRows(column)
			if err != nil {
				t.Error(err)
			}

			if countAfter != usersCountBefore {
				t.Errorf("Count users got = %d expect = %d", countAfter, usersCountBefore)
			}
		}
	}
}

func TestDatabase_UpdateUser(t *testing.T) {
	currentTime := time.Now()

	tests := []struct {
		name       string
		argsBefore structs.User
		argsAfter  structs.User
	}{
		{
			name: "Insert user",
			argsBefore: structs.User{
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
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   "",
				},
			},
			argsAfter: structs.User{
				ID: 1,
				Person: structs.Person{
					ID:       1,
					Name:     "Marco",
					Birthday: "01/01/2000",
					Gender:   "M",
				},
				CellNumber:  "10987654321",
				PhoneNumber: "0123456789",
				CPF:         "10987654321",
				Email:       "testmail@test.mail",
				Address: structs.Address{
					ID:           1,
					Number:       321,
					CEP:          "87654321",
					City:         "Curitiba",
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   "",
				},
				CreationDate: currentTime.Format("2006-01-02"),
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_userUpdate.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		_, err := dbDir.InsertUser(tt.argsBefore)
		if err != nil {
			t.Error(err)
			return
		}

		err = dbDir.UpdateUser(tt.argsAfter)
		if err != nil {
			t.Error(err)
			return
		}

		createdUser, err := dbDir.SelectUser(tt.argsAfter)
		if err != nil {
			t.Error(err)
			return
		}

		if createdUser != tt.argsAfter {
			gotJSON, _ := json.Marshal(createdUser)
			wantJSON, _ := json.Marshal(tt.argsAfter)

			t.Errorf("UpdateUser() got = %v, want %v", string(gotJSON), string(wantJSON))
		}
	}
}
