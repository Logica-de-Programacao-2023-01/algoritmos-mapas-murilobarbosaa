package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	chars := strings.Split(s, "")

	sort.Strings(chars)

	return strings.Join(chars, "")
}

func groupAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		sorted := sortString(word)
		anagramMap[sorted] = append(anagramMap[sorted], word)
	}

	return anagramMap
}

func main() {
	wordSlice := []string{"amor", "roma", "mora", "carro", "arroz"}

	anagramGroups := groupAnagrams(wordSlice)

	for _, words := range anagramGroups {
		fmt.Println(words)
	}
}