package structs

type Author struct {
	ID     int    `json:"id"`
	Person Person `json:"person"`
}
