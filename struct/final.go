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

func NewBook(id int, title, author string, genre string,pages int, price float64) (Book, error) {
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
	Genre string	`json:"genre"`
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

func describeGenre(genre string) string{
	switch genre{
	case "Fiction":
	return "Fantasy, Sci-Fi, Horror, Thriller"
	case "Non-Fiction":
	return "Business, Self-Help, Science"
	case "Technical":
	return "Programming, Engineering, Math"
	default:
		return "Uncategorized"
	}
}


func pageCategory(pages int) string{
	switch {
		case pages < 100:
			 return "Short Read"
		case pages < 300:
			 return "Medium Read"
		case pages < 600:
			 return "Long Read"
		case pages >= 600: 
			return "Epic Read"
	}
}

// Use your existing Base, Book, NewBook, loadShelf from Feature 5
func findBook(books []Book, title string) (Book, error) {
	for _, b := range books {
		if b.Title == title {
			return b, nil
		}
	}
	return Book{}, fmt.Errorf("book not found: %s", title)
}

func main() {
	books, err := loadShelf("shelf.json")
	if err != nil {
		fmt.Println("error loading json file", err)
		return
	}

	// Feature A — Price categorizer
	// Loop through all books
	// Use switch (no condition) to print price category:
	// > $70     → "💎 Premium: <title>"
	// $30-$70   → "📚 Standard: <title>"
	// < $30     → "💰 Budget: <title>"
	fmt.Println("\nPrice categorizer")
	for _, v := range books {
		switch {
		case v.Price > 70:
			fmt.Println("💎 Premium: ", v.Title)

		case v.Price >= 30 && v.Price <= 70:

			fmt.Println("📚 Standard: ", v.Title)
		case v.Price < 30:

			fmt.Println("💰 Budget: ", v.Title)
		}
	}

	// Feature C — Save with init statement
	// Add a new book to the shelf using NewBook
	// Save using saveShelf — call it with an IF INIT STATEMENT
	// Print "Saved!" on success, "Save failed: <err>" on failure
	fmt.Println("\nSave with init statement")

	b5, err := NewBook(4, "Who is this Allah", "Mosey Abdul", 800, 12.6)
	if err != nil {
		fmt.Println("Error creating book:", err)
		return
	}
	books = append(books, b5)

	if err := saveShelf(books, "shelf.json"); err != nil {
		fmt.Println("Save failed:", err)
	} else {
		fmt.Println("Saved!")
	}

	// Feature B — Search with init statement
	// Write a findBook(books []Book, title string) (Book, error)
	// function that returns a Book or error if not found
	// Call it using an IF INIT STATEMENT — no variables outside the if
	fmt.Println("\nSearch with init statement")
	if book, err := findBook(books, "Golang For Idiots"); err != nil {
		fmt.Println("Not found:", err)
	} else {
		fmt.Printf("Found: %s by %s\n", book.Title, book.Author)
	}

	// Feature A — Numbered list
	// Print each book with its position number:
	// "1. The Way Of The Water by Avatar Shin — $34.67"
	// "2. Geology For Idiots by Michael Bag — $60.67"
	// Use classic for loop with index
	fmt.Println("Numbered List")
	for index := 0; index < len(books); index++ {
		fmt.Printf("%d. %s by %s — $%.2f\n", index+1, books[index].Title, books[index].Author, books[index].Price)
	}

	//fmt.Println(books)

	// Feature B — Search with break
	// Search for a book by title "Geology For Idiots"
	// Print "Found: <title>" if found, "Not found" if not
	// Stop searching the moment you find it

	fmt.Println("\nSearch with break")
	found := false
	for _, v := range books {
		if v.Title == "Golang For Idiots" {
			fmt.Println("Found: ", v.Title)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Not Found")
	}

	// Feature C — Skip unavailable with continue
	// Set one book's InStock to false manually
	// Loop through books — skip books where InStock == false
	// Print only available books with prefix "Available: "
	fmt.Println("\nSkip unavailable with continue")
	books[1].InStock = false
	for _, k := range books {
		if !k.InStock {
			continue
		}
		fmt.Printf("Available: %s\n", k.Title)
	}
// Feature A — Genre dispatcher
// Add a Genre string field to your Book struct
// Update NewBook to accept genre string
// Write a function describeGenre(genre string) string that returns:
fmt.Println("\nGenre dispatcher")

// case "Fiction"    : 
// return "Fantasy, Sci-Fi, Horror, Thriller"
// case "Non-Fiction": 
// return "Business, Self-Help, Science"
// case "Technical"  : 
// return "Programming, Engineering, Math"
// case default      : 
// return "Uncategorized"
// Use switch with multiple values per case

// Feature B — fmt.Println("\nPage count category")
// Write pageCategory(pages int) string that returns:
fmt.Println("\nPage count category")
for _, s := range books{
	fmt.Printf("%s: %s\n", s.Title, s.Genre)
}
// case pages < 100: return "Short Read"
// case pages < 300: return "Medium Read"
// case pages < 600: return "Long Read"
// case pages >= 60: return "Epic Read"
// Use switch with no condition
// Loop through all books and print:
// "<title>: <category>"

// Feature C — Stock status dispatcher
// Loop through books
// Use switch on InStock (bool) to print:
// true  → "✅ In Stock: <title> — $<price>"
// false → "❌ Out of Stock: <title>"
// Set books[0].InStock = false to test both cases
fmt.Println("\nStock status dispatcher")
books[0].InStock = false
for _, b := range books{
	switch b.InStock{
	case true:
		fmt.Print("✅ In Stock: %s - %.2f\n", b.Title, b.Price)
	case false:
		fmt.Printf("❌ Out of Stock: %s\n", b.Title)
	}
}

}