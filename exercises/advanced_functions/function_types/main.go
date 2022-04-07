package main

import (
	"fmt"
)

// --------------------------------------------------
// The following code needs to be fixed and completed
// --------------------------------------------------

// Correct this function type to match printDetailed and printSummary
type transactionPrinter func() string

// Nothing needs to be changed in this method
func (tp transactionPrinter) printAllDelimited(transactions []Transaction, delimiter string) {
	fmt.Printf(delimiter)
	for _, t := range transactions {
		tp(t)
		fmt.Printf(delimiter)
	}
}

func (a *Account) Print() {
  // Complete this implementation so that printAllDelimited can be called on the
  // current value stored in a.printTransaction
}

// ---------------------------------------
// The following code should not be edited
// ---------------------------------------

type Account struct {
	transactions []Transaction
	printTransaction func(Transaction)
}

func (a *Account) setPrintMode(detailed bool) {
	if detailed {
		a.printTransaction = printDetailed
	} else {
		a.printTransaction = printSummary
	}
}

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
