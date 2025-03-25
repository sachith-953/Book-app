package main

import (
	"log"
	"net/http"
	"book-api/internal/book"
	"book-api/api"  
	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Register API routes
	r.HandleFunc("/books", book.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", book.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", book.GetBookById).Methods("GET")
	r.HandleFunc("/books/{id}", book.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", book.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		results := book.SearchBooks(query)
		api.WriteJSONResponse(w, http.StatusOK, results)
	}).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
