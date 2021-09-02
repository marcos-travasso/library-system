package structs

type Lending struct {
	ID         int          `json:"id"`
	User       User         `json:"user"`
	Book       Book         `json:"book"`
	LendDay    string       `json:"lendDay"`
	Returned   bool         `json:"returned"`
	Devolution []Devolution `json:"devolution"`
}

type Devolution struct {
	Date string `json:"date"`
}
