package book

import (
	"encoding/json"
	"net/http"

	"book-api/api" // Import the API response helper
	"book-api/internal/book/service" // Import service only
	"github.com/gorilla/mux"
)

// GetAllBooks returns all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := service.GetAllBooks() // Call service, NOT storage
	api.WriteJSONResponse(w, http.StatusOK, books)
}

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook api.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	createdBook := service.CreateBook(newBook) // Call service
	api.WriteJSONResponse(w, http.StatusCreated, createdBook)
}

// GetBookById retrieves a book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := service.GetBookById(params["id"]) // Call service
	if book == nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	api.WriteJSONResponse(w, http.StatusOK, book)
}

// UpdateBook updates a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedBook api.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	updated := service.UpdateBook(params["id"], updatedBook) // Call service
	if updated == nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	api.WriteJSONResponse(w, http.StatusOK, updated)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	success := service.DeleteBook(params["id"]) // Call service
	if !success {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
