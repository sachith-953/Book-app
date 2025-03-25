package service

import (
	"book-api/internal/book/storage"
	"book-api/api"
	"fmt"
	"time"
)

// GetAllBooks retrieves all books
func GetAllBooks() []api.Book {
	return storage.ReadBooks()
}

// CreateBook adds a new book
func CreateBook(newBook api.Book) api.Book {
	newBook.BookId = fmt.Sprintf("%d", time.Now().UnixNano()) // Generate a unique ID
	storage.WriteBook(newBook) // Call storage function
	return newBook
}

// GetBookById finds a book by ID
func GetBookById(id string) *api.Book {
	return storage.FindBookById(id) // Call storage function
}

// UpdateBook updates a book by ID
func UpdateBook(id string, updatedBook api.Book) *api.Book {
	return storage.UpdateBook(id, updatedBook) // Call storage function
}

// DeleteBook deletes a book by ID
func DeleteBook(id string) bool {
	return storage.DeleteBook(id) // Call storage function
}
