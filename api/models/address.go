package models

type Address struct {
	ID           int    `json:"id"`
	Number       int    `json:"number"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
}
