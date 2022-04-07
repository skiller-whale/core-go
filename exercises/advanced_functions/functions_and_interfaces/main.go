package main

import (
	"fmt"
)

// Interface for printing parts of a `Transaction`
type printer interface {
	print(t Transaction)
}

// Formatting functions. Each of these functions returns one part of a formatted `Transaction`.
func amountFormatter(t Transaction) string {
	return fmt.Sprintf("Amount transacted: %v", t.amount)
}

func typeFormatter(t Transaction) string {
	return fmt.Sprintf("Type of transaction: %v", t.creditOrDebit)
}

func separatorFormatter(t Transaction) string {
	return "---"
}

// Function type matching formatting functions and `printer` interface
type formatter func(Transaction) string

func (f formatter) print(t Transaction) {
	fmt.Println(f(t))
}

// Struct matching `printer` interface
type outputHighlight struct {
	credit string
	debit  string
}

func (oh outputHighlight) print(t Transaction) {
	if t.creditOrDebit == "credit" {
		fmt.Printf(oh.credit)
	} else {
		fmt.Printf(oh.debit)
	}
}

// Constants to set output colour
const Red string = "\033[31m"
const Green string = "\033[32m"
const Reset string = "\033[0m"

func main() {
	// You can use this value as an item inside `printers`
	oh := outputHighlight{
		credit: Green,
		debit:  Red,
	}

	printers := []printer{
		// Populate this slice with outputHighlight and formatting functions
	}

	for _, t := range transactions() {
		for _, p := range printers {
			p.print(t)
		}
	}
}

type Transaction struct {
	creditOrDebit string
	amount        int
	description   string
}

func transactions() []Transaction {
	return []Transaction{
		{
			creditOrDebit: "credit",
			amount:        560,
			description:   "Merchant Refund",
		},
		{
			creditOrDebit: "debit",
			amount:        240,
			description:   "Shopping",
		},
		{
			creditOrDebit: "debit",
			amount:        45,
			description:   "Internet service",
		},
		{
			creditOrDebit: "credit",
			amount:        100,
			description:   "Deposit",
		},
	}
}
