package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	s, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Current Balance from file :", s)
}

func writeBalanceToFile(balance float64) {
	b := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(b), 0644)
}

func main() {
	var accountBalance float64 = 1000
	writeBalanceToFile(accountBalance)

	for {

		fmt.Println("Wecome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Select your option: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice is :", choice)

		if choice == 1 {
			fmt.Println("Your balance is :", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter money to be deposited :")
			var depositAmount float64

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}

			accountBalance += depositAmount
			fmt.Println("Updated account Balance is :", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Print("Enter amount to be withdrawn :")

			var amountWithdrawn float64
			fmt.Scan(&amountWithdrawn)

			if amountWithdrawn > accountBalance {
				fmt.Println("Insufficient Amount entered!")
				fmt.Println("Current available balance is :", accountBalance)
			}

			accountBalance -= amountWithdrawn
			fmt.Println("Current account Balance is :", accountBalance)
			writeBalanceToFile(accountBalance)

		} else if choice == 4 {
			fmt.Println("Thank you for banking with us!")
			break
		} else {
			fmt.Print("Wrong choice")
			continue
		}

	}
}
