package models

type Lending struct {
	ID         int64      `json:"id"`
	User       User       `json:"user"`
	Book       Book       `json:"book"`
	LendDay    string     `json:"lendDay"`
	Returned   int        `json:"returned"`
	Devolution Devolution `json:"devolution"`
}
