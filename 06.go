package main

import "fmt"

func sumWordCounts(wordCountsList []map[string]int) map[string]int {
	sumCounts := make(map[string]int)

	for _, wordCounts := range wordCountsList {
		for word, count := range wordCounts {
			sumCounts[word] += count
		}
	}

	return sumCounts
}

func main() {
	wordCountslist := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 3, "go": 2, "world": 1},
		{"go": 1, "programming": 4},
	}

	sumCounts := sumWordCounts(wordCountslist)

	for word, count := range sumCounts {
		fmt.Printf("%s: %d \n", word, count)
	}
}
