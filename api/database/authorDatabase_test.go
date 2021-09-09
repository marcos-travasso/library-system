package database

import (
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
)

func TestDatabase_InsertAuthor(t *testing.T) {
	tests := []struct {
		name string
		args structs.Author
	}{
		{
			name: "First author",
			args: structs.Author{
				ID: 1,
				Person: structs.Person{
					ID:       1,
					Name:     "Clarice Lispector",
					Gender:   "F",
					Birthday: "10/12/1920",
				},
			},
		},
		{
			name: "Second author",
			args: structs.Author{
				ID: 2,
				Person: structs.Person{
					ID:       2,
					Name:     "Aldous Huxley",
					Gender:   "M",
					Birthday: "26/7/1894",
				},
			},
		},
		{
			name: "Repeated author",
			args: structs.Author{
				ID: 1,
				Person: structs.Person{
					ID:       1,
					Name:     "Clarice Lispector",
					Gender:   "F",
					Birthday: "10/12/1920",
				},
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_authorInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := dbDir.InsertAuthor(tt.args)
			if err != nil {
				t.Error(err)
			}

			if id != tt.args.ID {
				t.Errorf("Fail to check author id, got %d expect %d", id, tt.args.ID)
			}
		})
	}
}
