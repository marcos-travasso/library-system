package models

type Book struct {
	ID     int64  `json:"id"`
	Year   int    `json:"year"`
	Pages  int    `json:"pages"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}
