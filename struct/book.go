// This is the solution to BookShelf challenge

package main

import (
	"fmt"
)

// TODO 1: Define a Base struct with ID int, CreatedAt string, UpdatedAt string
type Base struct {
	ID        int
	CreatedAt string
	UpdatedAt string
}

// TODO 2: Define a Book struct that EMBEDS Base and has Title, Author, Pages, Price, InStock
type Book struct {
	Base
	Title   string
	Author  string
	Pages   int
	Price   float64
	InStock bool
}

// TODO 3: Update NewBook to accept id int as first argument
//
//	Set CreatedAt to "2024-01-01" and UpdatedAt to "2024-01-01" by default
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
func main() {
	// TODO 4: Create 3 books using updated NewBook
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
	// TODO 5: Print each book showing:
	//         ID, Title, Author, Price, CreatedAt
	//         e.g. "[1] Harry Potter by J.K. Rowling — $29.99 (created: 2024-01-01)"
	for _, book := range books {
		fmt.Printf("[%d] %s by %s - $%.2f (created: %s)\n", book.ID, book.Title, book.Author, book.Price, book.CreatedAt)
	}
	// TODO 6: Prove field promotion works —
	//         access b.ID and b.Base.ID and show they're equal
	fmt.Println("\nProving Field Promotion:")
	fmt.Println("b1.ID : ", b1.ID)
	fmt.Println("b1.Base.ID : ", b1.Base.ID)
	fmt.Printf("Are They Equal : %v\n", b1.ID == b1.Base.ID)
}
