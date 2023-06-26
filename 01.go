package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Olá, mundo!"
	wordCount := countWords(text)

	for word, count := range wordCount {
		fmt.Printf("%s: %d \n", word, count)
	}
}
