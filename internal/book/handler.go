package book

import (
	"encoding/json"
	"net/http"
	"book-api/api"
	"book-api/internal/book/service"
	"github.com/gorilla/mux"
)

// GetAllBooks returns all books in the system
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := service.GetAllBooks()
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
	book := service.CreateBook(newBook)
	api.WriteJSONResponse(w, http.StatusCreated, book)
}

// GetBookById retrieves a book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := service.GetBookById(params["id"])
	if book == nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	api.WriteJSONResponse(w, http.StatusOK, book)
}

// UpdateBook updates an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedBook api.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	book := service.UpdateBook(params["id"], updatedBook)
	if book == nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	api.WriteJSONResponse(w, http.StatusOK, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	success := service.DeleteBook(params["id"])
	if !success {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
