package database

import (
	"encoding/json"
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
	"time"
)

func TestInsertUser(t *testing.T) {
	tests := []struct {
		name    string
		args    structs.User
		want    int
		wantErr bool
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
			want:    1,
			wantErr: false,
		},
	}

	dbDir := Database{Dir: "./test_user_db.db"}
	dbDir.clearDatabase()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbDir.InsertUser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectUser(t *testing.T) {
	currentTime := time.Now()

	tests := []struct {
		name    string
		args    structs.User
		want    int
		wantErr bool
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
	}

	dbDir := Database{Dir: "./test_user_db.db"}
	dbDir.clearDatabase()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dbDir.InsertUser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v", err)
				return
			}

			got, err := dbDir.SelectUser(tt.args)
			if got != tt.args {
				gotJSON, _ := json.Marshal(got)
				wantJSON, _ := json.Marshal(tt.args)

				t.Errorf("SelectUser() got = %v, want %v", string(gotJSON), string(wantJSON))
			}
		})
	}
}
