package main

import "fmt"

func calculateExpenseEquelization(expenses map[string]float64) map[string]float64 {
	totalExpense := 0.0
	numPeople := len(expenses)

	for _, expense := range expenses {
		totalExpense += expense
	}

	equalizedExpenses := make(map[string]float64)
	for person, expense := range expenses {
		equalizedExpenses[person] = (totalExpense / float64(numPeople)) - expense
	}

	return equalizedExpenses
}

func main() {
	expenses := map[string]float64{
		"João":   100.0,
		"Maria":  50.0,
		"Pedro":  75.0,
		"Carlos": 120.0,
	}

	equelizedExpenses := calculateExpenseEquelization(expenses)

	for person, amount := range equelizedExpenses {
		fmt.Printf("%s: %.2f \n", person, amount)
	}
}
