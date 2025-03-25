package book

import (
	"encoding/json"
	"os"
	"log"
	"book-api/api"
)

const fileName = "books.json"

var books []api.Book // In-memory storage

// ReadBooks retrieves books from the JSON file
func ReadBooks() []api.Book {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening books file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&books)
	if err != nil {
		log.Println("Error decoding books:", err)
		return nil
	}
	return books
}

// WriteBook saves a new book
func WriteBook(newBook api.Book) {
	books = append(books, newBook)
	saveBooksToFile()
}

// FindBookById finds a book by its ID
func FindBookById(id string) *api.Book {
	for _, book := range books {
		if book.BookId == id {
			return &book
		}
	}
	return nil
}

// UpdateBook updates a book
func UpdateBook(id string, updatedBook api.Book) *api.Book {
	for i, book := range books {
		if book.BookId == id {
			books[i] = updatedBook
			saveBooksToFile()
			return &books[i]
		}
	}
	return nil
}

// DeleteBook deletes a book
func DeleteBook(id string) bool {
	for i, book := range books {
		if book.BookId == id {
			books = append(books[:i], books[i+1:]...)
			saveBooksToFile()
			return true
		}
	}
	return false
}

// saveBooksToFile writes books to the JSON file
func saveBooksToFile() {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating books file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(books)
	if err != nil {
		log.Println("Error encoding books:", err)
	}
}
