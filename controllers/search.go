package controllers

import (
	"strings"
	"github.com/gin-gonic/gin"
	"book-api/models"
)

// SearchBooks performs a case-insensitive search based on title and description
func SearchBooks(c *gin.Context) {
	keyword := c.DefaultQuery("q", "")
	if keyword == "" {
		c.JSON(400, gin.H{"error": "No search keyword provided"})
		return
	}

	// Read the books
	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Use goroutines and channels to improve search performance
	resultsChannel := make(chan models.Book)
	defer close(resultsChannel)

	// Split books into smaller chunks and search in parallel
	chunkSize := len(books) / 4 // Split the list into 4 chunks
	for i := 0; i < 4; i++ {
		go func(start, end int) {
			for _, book := range books[start:end] {
				if strings.Contains(strings.ToLower(book.Title), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(book.Description), strings.ToLower(keyword)) {
					resultsChannel <- book
				}
			}
		}(i*chunkSize, (i+1)*chunkSize)
	}

	// Collect the results from the channels
	var results []models.Book
	for i := 0; i < len(books)/chunkSize; i++ {
		result := <-resultsChannel
		results = append(results, result)
	}

	c.JSON(200, results)
}
