package models

type Author struct {
	ID     int64  `json:"id"`
	Person Person `json:"person"`
}
