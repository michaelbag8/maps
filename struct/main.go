package main

import (
	"fmt"
)

type Book struct {
	Title   string
	Author  string
	Pages   int
	Price   float64
	InStock bool
}

// This is a solution to a challenge
// TODO 1: validate title not empty
// TODO 2: validate pages > 0
// TODO 3: validate price >= 0
// TODO 4: return Book with InStock: true by default
func NewBook(title, author string, pages int, price float64) (Book, error) {

	if title == "" {
		return Book{}, fmt.Errorf("title can not be empty")
	}

	if pages < 0 {
		return Book{}, fmt.Errorf("pages can not be empty")
	}

	if pages <= 0 {
		return Book{}, fmt.Errorf("pages can not be empty")
	}

	return Book{
		Title:   title,
		Author:  author,
		Pages:   pages,
		Price:   price,
		InStock: true,
	}, nil
}

// TODO 5: Create 4 books using NewBook
// TODO 6: Store in []Book
// TODO 7: Print all books formatted:
// TODO 8: Filter and print only books under $25
// TODO 9: Find and print the most expensive book
func main() {

	books := []Book{}
	b1, err := NewBook("Why", "Bag", 300, 2.45)
	if err == nil {
		books = append(books, b1)
	}

	b2, err := NewBook("Why now", "Michael", 600, 76.6)
	if err == nil {
		books = append(books, b2)
	}

	b3, err := NewBook("Why me", "Mike", 700, 23.8)
	if err == nil {
		books = append(books, b3)
	}

	b4, err := NewBook("Why Not You", "Zoe", 900, 24.7)
	if err == nil {
		books = append(books, b4)
	}

	fmt.Println("\nAll Books Formatted")
	for _, b := range books {
		fmt.Printf("%s by %s - %d pages, $%v\n",
			b.Title, b.Author, b.Pages, b.Price)
	}

	fmt.Println("\nBooks under $25:")
	for _, c := range books {
		if c.Price < 25 {
			fmt.Printf("%s by %s - %d pages, $%.2f\n",
				c.Title, c.Author, c.Pages, c.Price)

		}
	}

	if len(books) > 0 {
		max := books[0]
		for _, d := range books {
			if d.Price > max.Price {
				max = d
			}

		}
		fmt.Printf("\nMost expensive: %s by %s — $%.2f\n",
			max.Title, max.Author, max.Price)
	}
}
