package book

import (
	"book-api/api"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"book-api/internal/book/storage"
)

var books []api.Book

// Initialize by reading from the JSON file
func init() {
	err := storage.ReadBooksFromFile()
	if err != nil {
		log.Println("Error reading books data:", err)
	}
}

// GetAllBooks returns all books in the system
func GetAllBooks() []api.Book {
	return books
}

// CreateBook adds a new book to the list
func CreateBook(newBook api.Book) api.Book {
	newBook.BookId = fmt.Sprintf("%s", time.Now().UnixNano())
	books = append(books, newBook)
	storage.WriteBooksToFile()
	return newBook
}

// GetBookById finds a book by its ID
func GetBookById(id string) *api.Book {
	for _, book := range books {
		if book.BookId == id {
			return &book
		}
	}
	return nil
}

// UpdateBook updates a book by its ID
func UpdateBook(id string, updatedBook api.Book) *api.Book {
	for index, book := range books {
		if book.BookId == id {
			books[index] = updatedBook
			storage.WriteBooksToFile()
			return &books[index]
		}
	}
	return nil
}

// DeleteBook deletes a book by its ID
func DeleteBook(id string) bool {
	for index, book := range books {
		if book.BookId == id {
			books = append(books[:index], books[index+1:]...)
			storage.WriteBooksToFile()
			return true
		}
	}
	return false
}

// SearchBooks performs a keyword search on book titles and descriptions (optimized with concurrency)
func SearchBooks(query string) []api.Book {
	var results []api.Book
	ch := make(chan []api.Book)
	
	// Split the list into two parts and search concurrently
	go searchBooksInSubset(books[:len(books)/2], query, ch)
	go searchBooksInSubset(books[len(books)/2:], query, ch)
	
	// Collect results from goroutines
	for i := 0; i < 2; i++ {
		results = append(results, <-ch...)
	}
	return results
}

func searchBooksInSubset(books []api.Book, query string, ch chan []api.Book) {
	var result []api.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) || 
			strings.Contains(strings.ToLower(book.Description), strings.ToLower(query)) {
			result = append(result, book)
		}
	}
	ch <- result
}
