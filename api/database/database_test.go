package database

import (
	"github.com/marcos-travasso/library-system/api"
	"os"
	"testing"
)

var db = Database{
	"./test_db.db",
}

func TestDatabase_CreateDatabase(t *testing.T) {
	db.CreateDatabase()

	if _, err := os.Stat(db.Dir); os.IsNotExist(err) {
		t.Errorf("Fail to create database file")
	}
}

func TestDatabase_insertPersonStatement(t *testing.T) {
	tests := []struct {
		name   string
		fields Database
		args   api.Person
		want   string
	}{
		{name: "One person",
			fields: db,
			args: api.Person{
				Name:     "Marcos",
				Birthday: "01/01/2000",
				Gender:   "M",
			},
			want: "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"01/01/2000\")",
		},
		{name: "Empty person",
			fields: db,
			args: api.Person{
				Name:     "",
				Birthday: "",
				Gender:   "",
			},
			want: "",
		}, {name: "Empty name",
			fields: db,
			args: api.Person{
				Name:     "",
				Birthday: "01/01/2000",
				Gender:   "F",
			},
			want: "",
		}, {name: "Empty birthday",
			fields: db,
			args: api.Person{
				Name:     "Marcos",
				Birthday: "",
				Gender:   "M",
			},
			want: "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				Dir: tt.fields.Dir,
			}
			if got := db.insertPersonStatement(tt.args); got != tt.want {
				t.Errorf("insertPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
