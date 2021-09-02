package structs

type User struct {
	ID          int     `json:"id"`
	Person      Person  `json:"person"`
	CellNumber  string  `json:"cellNumber"`
	PhoneNumber string  `json:"phoneNumber"`
	CPF         string  `json:"cpf"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
}

type Address struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
}
