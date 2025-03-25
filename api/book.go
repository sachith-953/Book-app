package api

import "time"

// Book struct defines the structure of the Book entity
type Book struct {
	BookId          string    `json:"bookId"`
	AuthorId        string    `json:"authorId"`
	PublisherId     string    `json:"publisherId"`
	Title           string    `json:"title"`
	PublicationDate time.Time `json:"publicationDate"`
	ISBN            string    `json:"isbn"`
	Pages           int       `json:"pages"`
	Genre           string    `json:"genre"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Quantity        int       `json:"quantity"`
}
