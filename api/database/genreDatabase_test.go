package database

import (
	"github.com/marcos-travasso/library-system/api/structs"
	"testing"
)

func TestDatabase_insertGenre(t *testing.T) {
	tests := []struct {
		name string
		args structs.Genre
	}{
		{
			name: "First genre",
			args: structs.Genre{
				ID:   1,
				Name: "Romance",
			},
		},
		{
			name: "Second genre",
			args: structs.Genre{
				ID:   2,
				Name: "Tale",
			},
		},
		{
			name: "Repeated genre",
			args: structs.Genre{
				ID:   1,
				Name: "Romance",
			},
		},
	}

	dbDir := Database{Dir: "./temp/test_genreInsert.db"}
	err := dbDir.clearDatabase()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := dbDir.insertGenre(tt.args)
			if err != nil {
				t.Error(err)
			}

			if id != tt.args.ID {
				t.Errorf("Fail to check genre id, got %d expect %d", id, tt.args.ID)
			}
		})
	}
}
