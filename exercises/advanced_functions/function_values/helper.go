package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	date          time.Time
	creditOrDebit string
	amount        int
	description   string
}

var account = Account{
	transactions: []Transaction{
		{
			date:          time.Date(2022, time.February, 22, 22, 22, 22, 22, time.UTC),
			creditOrDebit: "credit",
			amount:        560,
			description:   "Merchant Refund",
		},
		{
			date:          time.Date(2022, time.February, 25, 20, 12, 2, 32, time.UTC),
			creditOrDebit: "debit",
			amount:        45,
			description:   "Amazon Shopping",
		},
	},
	printTransaction: printSummary,
}

func (a *Account) Print() {
	fmt.Println("---")
	for _, t := range a.transactions {
		a.printTransaction(t)
		fmt.Println("---")
	}
}

func printInstructions(detailed bool) {
	fmt.Println()
	fmt.Println("-----------------------------")
	fmt.Println("Enter an option:")
	fmt.Println("1 - Print transactions")
	if detailed {
		fmt.Println("2 - Toggle print mode. Current mode is detailed.")
	} else {
		fmt.Println("2 - Toggle print mode. Current mode is summary.")
	}
	fmt.Println("    (Any other input to exit)")
	fmt.Println()
}

func getUserInput() int {
	var n int
	fmt.Scanln(&n)
	return n
}
