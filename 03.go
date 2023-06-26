package main

import "fmt"

func sumMapsValues(m map[string]int) int {
	sum := 0

	for _, value := range m {
		sum += value
	}

	return sum
}

func main() {
	m := map[string]int{
		"valor1": 10,
		"valor2": 20,
		"valor3": 30,
	}

	total := sumMapsValues(m)
	fmt.Println("Soma dos valores:", total)
}
