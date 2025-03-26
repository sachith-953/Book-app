package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"book-api/models"
)

// Helper function to read books from a JSON file
func readBooksFromFile() ([]models.Book, error) {
	file, err := os.Open("books.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []models.Book
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&books); err != nil {
		return nil, err
	}

	return books, nil
}

// Helper function to write books to a JSON file
func writeBooksToFile(books []models.Book) error {
	file, err := os.Create("books.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(books)
}

// GetBooks returns a list of all books
func GetBooks(c *gin.Context) {
	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook adds a new book
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID
	newBook.BookId = time.Now().String()

	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)

	if err := writeBooksToFile(books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

// GetBook returns a single book by ID
func GetBook(c *gin.Context) {
	id := c.Param("id")
	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, book := range books {
		if book.BookId == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// UpdateBook updates a book by ID
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.BookId == id {
			books[i] = updatedBook
			if err := writeBooksToFile(books); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.BookId == id {
			books = append(books[:i], books[i+1:]...)
			if err := writeBooksToFile(books); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
