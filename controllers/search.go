package controllers

import (
	"book-api/models"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// SearchBooks performs a case-insensitive search based on title and description
func SearchBooks(c *gin.Context) {
	keyword := c.DefaultQuery("q", "")
	if keyword == "" {
		c.JSON(400, gin.H{"error": "No search keyword provided"})
		return
	}

	books, err := readBooksFromFile()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var results []models.Book
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := (len(books) + 3) / 4 // Safe chunk size (avoid losing books)

	for i := 0; i < len(books); i += chunkSize {
		end := i + chunkSize
		if end > len(books) {
			end = len(books)
		}

		wg.Add(1)
		go func(chunk []models.Book) {
			defer wg.Done()
			var localResults []models.Book
			for _, book := range chunk {
				if strings.Contains(strings.ToLower(book.Title), strings.ToLower(keyword)) ||
					strings.Contains(strings.ToLower(book.Description), strings.ToLower(keyword)) {
					localResults = append(localResults, book)
				}
			}
			mu.Lock()
			results = append(results, localResults...)
			mu.Unlock()
		}(books[i:end])
	}

	wg.Wait()

	c.JSON(200, results)
}
