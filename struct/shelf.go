package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO 1: Add JSON tags to Base struct
type Base struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func NewBook(id int, title, author string, pages int, price float64) (Book, error) {
	if title == "" {
		return Book{}, fmt.Errorf("title can not be empty")
	}
	if price < 0 {
		return Book{}, fmt.Errorf("price cannot be negative, got %.2f", price)
	}
	if pages <= 0 {
		return Book{}, fmt.Errorf("pages can not be empty")
	}
	return Book{
		Base:    Base{ID: id, CreatedAt: "2024-01-01", UpdatedAt: "2024-01-01"},
		Title:   title,
		Author:  author,
		Pages:   pages,
		Price:   price,
		InStock: true,
	}, nil
}

// TODO 2: Add JSON tags to Book struct
//   - InStock should be omitted if false
//   - UpdatedAt should be omitted if empty
type Book struct {
	Base
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Pages   int     `json:"pages"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock,omitempty"`
}

// TODO 3: Write saveShelf(books []Book, filename string) error
//
//	that marshals books to JSON and saves to file
func saveShelf(books []Book, filename string) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal books: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil

}

// TODO 4: Write loadShelf(filename string) ([]Book, error)
//
//	that reads JSON file and unmarshals into []Book
func loadShelf(filename string) ([]Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	var books []Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("failed to unmarshal books: %w", err)
	}
	return books, nil
}

func main() {
	// TODO 5: Create 3 books using NewBook
	books := []Book{}
	b1, err := NewBook(1, "The Way Of The Water", "Avatar Shin", 200, 34.67)
	if err != nil {
		fmt.Println("error: creating a book")
	} else {
		books = append(books, b1)
	}
	b2, err := NewBook(2, "Geology For Idiots", "Michael Bag", 300, 60.67)
	if err != nil {
		fmt.Println("error: creating a book")
	} else {
		books = append(books, b2)
	}
	b3, err := NewBook(3, "Golang For Idiots", "Aboh Oche", 600, 98.67)
	if err != nil {
		fmt.Println("error: creating a book")
	} else {
		books = append(books, b3)
	}
	// TODO 6: Save to "shelf.json" using saveShelf
	err = saveShelf(books, "shelf.json")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// TODO 7: Load back from "shelf.json" using loadShelf
	data, err := loadShelf("shelf.json")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// TODO 8: Print loaded books to prove round-trip works
	//         "[1] Harry Potter by J.K. Rowling — $29.99"

	for _, book := range data {
		fmt.Printf("[%d] %s by %s - $%.2f (created: %s)\n", book.ID, book.Title, book.Author, book.Price, book.CreatedAt)
	}
}

