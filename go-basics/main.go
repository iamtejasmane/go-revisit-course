package main

import (
	"fmt"
)

func main() {
	revenue := userInput("Enter Revenue: ")
	expenses := userInput("Enter Expenses: ")
	tax_rate := userInput("Enter Tax Rate: ")

	ebt, profit, ration := calculateEBT(revenue, expenses, tax_rate)
	fmt.Println(ebt, profit, ration)

}

func calculateEBT(revenue, expenses, tax_rate float64) (ebt, profit, ration float64) {
	// EBT = Revenue + Tax
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ration = ebt / profit

	return ebt, profit, ration
}

func userInput(text string) (userValue float64) {
	fmt.Print(text)
	fmt.Scan(&userValue)

	return userValue
}
