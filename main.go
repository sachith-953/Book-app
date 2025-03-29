package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

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

// Custom time layout for the publicationDate
const dateLayout = "2006-01-02"

// Function to read the JSON data from a file
func readBooksFromFile() ([]Book, error) {
	file, err := os.Open("books.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var books []Book
	if err := json.Unmarshal(bytes, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// Function to write the JSON data to a file
func writeBooksToFile(books []Book) error {
	bytes, err := json.Marshal(books)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("books.json", bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to handle unmarshaling the Book entity with custom time parsing
func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		PublicationDate string `json:"publicationDate"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the publication date with custom layout
	parsedDate, err := time.Parse(dateLayout, aux.PublicationDate)
	if err != nil {
		return fmt.Errorf("parsing time %s as %s: %v", aux.PublicationDate, dateLayout, err)
	}
	b.PublicationDate = parsedDate
	return nil
}

// Function to handle searching books based on title and description
func searchBooks(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		http.Error(w, "No search keyword provided", http.StatusBadRequest)
		return
	}

	// Read all books from file
	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	var results []Book
	for _, book := range books {
		// Check if the keyword exists in either title or description (case-insensitive)
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(book.Description), strings.ToLower(keyword)) {
			results = append(results, book)
		}
	}

	if len(results) == 0 {
		http.Error(w, "No books found matching your search", http.StatusNotFound)
		return
	}

	// Return the found books as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// Function to get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Function to create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	// Add the new book to the list
	books = append(books, newBook)

	// Write the updated list back to the file
	err = writeBooksToFile(books)
	if err != nil {
		http.Error(w, "Could not save books data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// Function to get a single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]

	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		if book.BookId == bookId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Function to update a book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]

	var updatedBook Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	for i, book := range books {
		if book.BookId == bookId {
			// Update the book details
			books[i] = updatedBook
			err = writeBooksToFile(books)
			if err != nil {
				http.Error(w, "Could not save books data", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Function to delete a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]

	books, err := readBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read books data", http.StatusInternalServerError)
		return
	}

	for i, book := range books {
		if book.BookId == bookId {
			// Remove the book from the slice
			books = append(books[:i], books[i+1:]...)
			err = writeBooksToFile(books)
			if err != nil {
				http.Error(w, "Could not save books data", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Main function to set up the routes and start the server
func main() {
	r := mux.NewRouter()

	// Root endpoint displaying welcome message and available endpoints
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Welcome message with the available endpoints
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to the Book API!\n")
		fmt.Fprintf(w, "Here are the available endpoints:\n")
		fmt.Fprintf(w, "- GET /books         - Get a list of all books\n")
		fmt.Fprintf(w, "- POST /books        - Create a new book\n")
		fmt.Fprintf(w, "- GET /books/{id}    - Get a single book by ID\n")
		fmt.Fprintf(w, "- PUT /books/{id}    - Update a book by ID\n")
		fmt.Fprintf(w, "- DELETE /books/{id} - Delete a book by ID\n")
		fmt.Fprintf(w, "- GET /search        - Search books by title/description (use query parameter ?q=your_search_term)\n\n")
		fmt.Fprintf(w, "Example:\n")
		fmt.Fprintf(w, "GET /search?q=great (Enter Book Title) \n")
	}).Methods("GET")

	// Other existing routes for books and search
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Add the search route
	r.HandleFunc("/search", searchBooks).Methods("GET")

	log.Println("Server starting on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
