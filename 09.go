package main

import "fmt"

func fibonacciSequence(n int) map[int]int {
	sequence := make(map[int]int)

	sequence[0] = 0
	sequence[1] = 1

	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}

func main() {
	n := 10

	fibSeq := fibonacciSequence(n)

	for i, num := range fibSeq {
		fmt.Printf("%d: %d \n", i, num)
	}
}
