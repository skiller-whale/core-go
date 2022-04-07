package main

import (
	"fmt"
)

type Account struct {
	transactions []Transaction
	// Add a new field
}

func (a *Account) setPrintMode(detailed bool) {
	// Implement this function
}

// DO NOT EDIT BELOW THIS

func printDetailed(t Transaction) {
	fmt.Println("Date: ", t.date.Format("2006-January-02"))
	fmt.Println("Amount transacted: ", t.amount)
	fmt.Println("Type of transaction: ", t.creditOrDebit)
}

func printSummary(t Transaction) {
	fmt.Printf("%v was %ved on %v\n", t.amount, t.creditOrDebit, t.date.Format("2006-January-02"))
}

func main() {
	var detailed = false
	for true {
		printInstructions(detailed)
		switch getUserInput() {
		case 1:
			account.Print()
		case 2:
			detailed = !detailed
			account.setPrintMode(detailed)
		default:
			return
		}
	}
}
