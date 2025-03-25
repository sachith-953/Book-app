package book

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
)

const fileName = "books.json"

// ReadBooksFromFile reads the books from the JSON file
func ReadBooksFromFile() error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		return err
	}

	return nil
}

// WriteBooksToFile writes the books data to the JSON file
func WriteBooksToFile() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(byteValue)
	if err != nil {
		return err
	}

	return nil
}
