package main

import (
	"fmt"
	"strings"
)

func countLettersInWords(phrase string) map[string]map[rune]int {
	wordCounts := make(map[string]map[rune]int)

	words := splitIntoWords(phrase)

	for _, word := range words {
		letterCounts := make(map[rune]int)

		for _, letter := range word {
			letterCounts[letter]++
		}

		wordCounts[word] = letterCounts
	}

	return wordCounts
}

func splitIntoWords(phrase string) []string {
	return strings.Fields(phrase)
}

func main() {
	phrase := "Hello, world! Welcome to Brasil."

	wordLetterCounts := countLettersInWords(phrase)

	for word, letterCounts := range wordLetterCounts {
		fmt.Printf("%s: %v \n", word, letterCounts)
	}
}
