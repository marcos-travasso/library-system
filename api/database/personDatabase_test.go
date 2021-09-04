package database

import (
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
)

func TestDatabase_InsertPerson(t *testing.T) {
	tests := []struct {
		name    string
		args    structs.Person
		want    int
		wantErr bool
	}{
		{
			name: "First person",
			args: structs.Person{
				ID:       0,
				Name:     "Marcos",
				Birthday: "01/01/2002",
				Gender:   "M",
			},
			want: 1,
		},
		{
			name: "Empty person",
			args: structs.Person{
				ID:       0,
				Name:     "",
				Birthday: "",
				Gender:   "",
			},
			wantErr: true,
		},
		{
			name: "Second person",
			args: structs.Person{
				ID:       0,
				Name:     "Clarice",
				Birthday: "",
				Gender:   "",
			},
			want: 2,
		},
	}

	dbDir := Database{Dir: "./test_person_db.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := dbDir.InsertPerson(tt.args)
			if tt.wantErr && err != nil {
				return
			} else if tt.wantErr && err == nil {
				t.Error(err)
			}

			if id != tt.want {
				t.Error(err)
			}
		})
	}
}
