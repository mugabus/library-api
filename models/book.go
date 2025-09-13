package models

type Book struct {
	BookID   int    `json:"book_id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Borrowed bool   `json:"borrowed"`
}
