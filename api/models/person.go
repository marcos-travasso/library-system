package models

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
