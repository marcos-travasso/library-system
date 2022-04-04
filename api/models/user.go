package models

type User struct {
	ID           int64   `json:"id"`
	Person       Person  `json:"person"`
	CellNumber   string  `json:"cellNumber"`
	PhoneNumber  string  `json:"phoneNumber"`
	CPF          string  `json:"cpf"`
	Email        string  `json:"email"`
	Address      Address `json:"address"`
	Responsible  Person  `json:"responsible"`
	CreationDate string  `json:"creationDate"`
}
