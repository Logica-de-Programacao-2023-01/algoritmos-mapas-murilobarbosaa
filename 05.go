package main

import "fmt"

func charFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)

	for _, char := range s {
		frequency[char]++
	}

	return frequency
}

func main() {
	str := "Hello, World"
	freqMap := charFrequency(str)

	for char, count := range freqMap {
		fmt.Printf("%c: %d \n", char, count)
	}
}
