package main

import (
	"fmt"
	"strings"
)

// MapCSVToIndex maps comma-separated values to their respective positions.
func MapCSVToIndex(board string) map[int]string {
	result := make(map[int]string)
	parts := strings.Split(board, ",")
	for i, v := range parts {
		result[i] = v
	}
	return result
}

// MapWordsToIndex maps whitespace-separated words to their numeric indices.
func MapWordsToIndex(str string) map[int]string {
	blocks := make(map[int]string)
	words := strings.Fields(str)
	for k, v := range words {
		blocks[k] = v
	}
	return blocks
}

// MapWordLastOccurrence stores each unique word with the index of its last appearance.
func MapWordLastOccurrence(text string) map[string]int {
	words := strings.Fields(text)
	blocks := map[string]int{}
	for i, value := range words {
		blocks[value] = i
	}
	return blocks
}

// ChunkWordsIntoTriplets groups a string of words into slices of three.
func ChunkWordsIntoTriplets(text string) (blocks [][]string) {
	words := strings.Fields(text)
	n := len(words)
	groupSize := 3

	for i := 0; i < n; i += groupSize {
		end := i + groupSize
		if end > n {
			end = n
		}
		blocks = append(blocks, words[i:end])
	}
	return
}

func main() {
	words := "welcome to abuja where things are happening"
	fmt.Println(MapWordsToIndex(words))

	data := "apple orange apple banana orange apple banana apple banana"
	fmt.Println(MapWordLastOccurrence(data))

	fmt.Println("------------------------------")
	fmt.Println(ChunkWordsIntoTriplets(data))

	boards := "X,_,O,_,X,_,O,_,X"
	fmt.Println(MapCSVToIndex(boards))

	// fixquote
	fmt.Println(fixQuote(" i have a book . . . that i love reading , everyday"))

	text := "TODO make this bad code better"
	ab := "this code has been fix"
	abs := "pending make this bad code better"
	abc := "fail make this bad code better"
	fmt.Println(replaceWord(text))
	fmt.Println(replaceWord(ab))
	fmt.Println(replaceWord(abs))
	fmt.Println(replaceWord(abc))

	fmt.Println("---------------")
	err := printChar("BAG")
	if err != nil {
		fmt.Println("Error printing character")
	}

	fmt.Println(fixArticle("a 'apple', an goat, A orange, An banana"))
}
