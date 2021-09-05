package database

import (
	"database/sql"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
	"testing"
)

var createdPersonIDs = 0

func TestDatabase_InsertPerson(t *testing.T) {
	tests := []struct {
		name    string
		args    structs.Person
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
			createdPersonIDs++

			if id != createdPersonIDs {
				t.Error(err)
			}
		})
	}
}

func TestDatabase_getLastPerson(t *testing.T) {
	dbDir := Database{Dir: "./test_person_db.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		return
	}

	var db = initializeDatabase(dbDir)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error to close database: %v", err)
		}
	}(db)

	id, err := getLastPersonID(db)
	if err != nil {
		t.Error(err)
	}

	if id > createdPersonIDs {
		t.Error("id error")
	}
}