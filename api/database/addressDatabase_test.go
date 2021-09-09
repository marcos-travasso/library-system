package database

import (
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
)

func TestDatabase_InsertAddress(t *testing.T) {
	tests := []struct {
		name string
		args structs.Address
		want int
	}{
		{
			name: "First address",
			args: structs.Address{
				ID:           0,
				Number:       123,
				CEP:          "12345678",
				City:         "Curitiba",
				Neighborhood: "Centro",
				Street:       "XV",
				Complement:   "",
			},
			want: 1,
		},
	}

	dbDir := Database{Dir: "./temp/test_addressInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbDir.insertAddress(tt.args)
			if err != nil {
				t.Errorf("InsertAddress() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("InsertAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_DeleteAddress(t *testing.T) {
	{
		tests := []struct {
			name string
			args structs.Address
		}{
			{
				name: "First address",
				args: structs.Address{
					ID:           1,
					Number:       123,
					CEP:          "12345678",
					City:         "Curitiba",
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   "",
				},
			},
			{
				name: "Second address",
				args: structs.Address{
					ID:           2,
					Number:       321,
					CEP:          "98765432",
					City:         "Curitiba",
					Neighborhood: "Centro",
					Street:       "XV",
					Complement:   "",
				},
			},
			{
				name: "Third address",
				args: structs.Address{
					ID:           3,
					Number:       0,
					CEP:          "123456789",
					City:         "",
					Neighborhood: "",
					Street:       "",
					Complement:   "",
				},
			},
		}

		dbDir := Database{Dir: "./temp/test_addressDelete.db"}
		err := dbDir.clearDatabase()
		if err != nil {
			t.Error(err)
			return
		}

		for _, tt := range tests {
			_, err := dbDir.insertAddress(tt.args)
			if err != nil {
				t.Error(err)
			}
		}

		addressCountBefore, err := dbDir.countRows("Enderecos")
		if err != nil {
			t.Error(err)
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := dbDir.deleteAddress(tt.args)
				if err != nil {
					t.Error(err)
				}
				addressCountBefore--
			})
		}

		columnsToCount := []string{"Enderecos"}
		for _, column := range columnsToCount {
			countAfter, err := dbDir.countRows(column)
			if err != nil {
				t.Error(err)
			}

			if countAfter != addressCountBefore {
				t.Errorf("Count address got = %d  expect = %d", countAfter, addressCountBefore)
			}
		}
	}
}
