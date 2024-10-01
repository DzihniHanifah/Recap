package main

import (
	"fmt"
)

func main() {
	const months = 12
	var expenses [months]float64

	//Input monthly expenses
	for i := 0; i < months; i++ {
		fmt.Printf("Enter the expenses for month %d: ", i+1)
		fmt.Scan(&expenses[i])
	}
	//Calculate total and average expenses
	var total float64
	for _, expenses := range expenses {
		total += expenses
	}
	average := total / months

	fmt.Printf("Total expenses for the year: %.2f\n", total)
	fmt.Printf("Average monthly expenses: %.2f\n", average)

}
