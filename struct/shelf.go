
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// TODO 1: Add JSON tags to Base struct
type Base struct {
	ID        int `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
// TODO 2: Add JSON tags to Book struct
//         - InStock should be omitted if false
//         - UpdatedAt should be omitted if empty
type Book struct {
	Base
	Title   string `json:"title"`
	Author  string `json:"author"`
	Pages   int		`json:"pages"`
	Price   float64 `json:"price"`
	InStock bool 	`json:"in_stock,omitempty"`
}


// TODO 3: Write saveShelf(books []Book, filename string) error
//         that marshals books to JSON and saves to file
func saveShelf(books []Book, filename string) error{

}

// TODO 4: Write loadShelf(filename string) ([]Book, error)
//         that reads JSON file and unmarshals into []Book
func loadShelf(filename string) ([]Book, error){
	
}

func main() {
    // TODO 5: Create 3 books using NewBook
    // TODO 6: Save to "shelf.json" using saveShelf
    // TODO 7: Load back from "shelf.json" using loadShelf
    // TODO 8: Print loaded books to prove round-trip works
    //         "[1] Harry Potter by J.K. Rowling — $29.99"
}