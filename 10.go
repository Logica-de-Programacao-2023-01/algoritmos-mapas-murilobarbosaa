package main

import "fmt"

func countPairFrequencies(nums []int) map[[2]int]int {
	pairFrequencies := make(map[[2]int]int)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			pair := [2]int{nums[i], nums[j]}
			pairFrequencies[pair]++
		}
	}

	return pairFrequencies
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 4, 5, 4, 5}

	pairFreq := countPairFrequencies(nums)

	for pair, freq := range pairFreq {
		fmt.Printf("%v: %d \n", pair, freq)
	}
}
