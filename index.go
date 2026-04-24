package main

import (
	"fmt"
	"strings"
)

// MultiPunctuations fix muiltple punctuations
func MultiPunctuations(str string) string {
	replacement := map[string]string{
		". . .": "...",
		"! ! !": "!!!",
		"! ?":   "!?",
		": :":   "::",
	}
	for oldVal, newVal := range replacement {
		str = strings.ReplaceAll(str, oldVal, newVal)
	}
	return str
}

// fixquote fix punctuations
func fixQuote(str string) string {
	str = MultiPunctuations(str)
	punctuations := []string{",", ".", ":", "?", "'", "!", "-", "_", "\""}
	for _, p := range punctuations {
		str = strings.ReplaceAll(str, " "+p, p)
	}
	return str
}

// reaplaceWord replaces words in a text to the matched ones from the replacement map
func replaceWord(text string) string {
	texts := strings.Fields(text)
	replacements := map[string]string{
		"TODO":    "DONE",
		"bad":     "good",
		"error":   "success",
		"fix":     "resolved",
		"missing": "found",
		"pending": "completed",
		"fail":    "pass",
	}

	for i, value := range texts {
		if newWord, ok := replacements[value]; ok {
			texts[i] = newWord
		}
	}

	return strings.Join(texts, " ")

}

// printChar print characters
func printChar(text string) error {

	characters := map[rune][]string{
		'A': {
			"   /\\   ",
			"  /  \\  ",
			" / /\\ \\ ",
			"/ ____ \\",
			"| |  | |",
			"|_|  |_|",
		},

		'B': {
			"|~~~\\  ",
			"| |\\ \\ ",
			"| |_/ /",
			"|  _ < ",
			"| |_/ /",
			"|____/ ",
		},

		'G': {
			"  ____ ",
			" / ___|",
			"| | __ ",
			"| ||_ |",
			"| |__| |",
			" \\____/ ",
		},
	}

	for row := 0; row < 6; row++ {
		for _, ch := range text {
			fmt.Print(characters[ch][row])
		}
		fmt.Println()
	}
	return nil
}
func fixArticle(text string) string {
	words := strings.Fields(text)

	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for i := 0; i < len(words)-1; i++ {
		article := words[i]

		if article != "a" && article != "an" && article != "A" && article != "An" {
			continue
		}

		nextIndex := i + 1

		
		if words[nextIndex] == "'" {
			nextIndex++
		}

		
		if nextIndex >= len(words) {
			continue
		}

		nextWord := words[nextIndex]
		nextWord = strings.Trim(nextWord, "'\"")

		if nextWord == "" {
			continue
		}

		firstLetter := []rune(strings.ToLower(nextWord))[0]

		if vowels[firstLetter] {
			switch article {
			case "a":
				words[i] = "an"
			case "A":
				words[i] = "An"
			}
		} else {
			switch article {
			case "an":
				words[i] = "a"
			case "An":
				words[i] = "A"
			}
		}
	}

	return strings.Join(words, " ")
}