package api

import "testing"

func TestPerson_insertPersonStatement(t *testing.T) {
	tests := []struct {
		name string
		args Person
		want string
	}{
		{name: "One person",
			args: Person{
				Name:     "Marcos",
				Birthday: "01/01/2000",
				Gender:   "M",
			},
			want: "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"01/01/2000\")",
		},
		{name: "Empty person",
			args: Person{
				Name:     "",
				Birthday: "",
				Gender:   "",
			},
			want: "",
		}, {name: "Empty name",
			args: Person{
				Name:     "",
				Birthday: "01/01/2000",
				Gender:   "F",
			},
			want: "",
		}, {name: "Empty birthday",
			args: Person{
				Name:     "Marcos",
				Birthday: "",
				Gender:   "M",
			},
			want: "INSERT INTO Pessoas(Nome, Genero, Nascimento) values (\"Marcos\", \"M\", \"\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.insertPersonStatement(); got != tt.want {
				t.Errorf("insertPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
