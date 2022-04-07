package main

import (
	"errors"
	"fmt"
)

func (a Account) runningBalanceCalculator(balance int, i int) func() (int, error) {
	// Implement this method using a closure
}

// This method returns a new balance based on the amount and type of `transaction`.
// You do not need to change this function.
func adjustBalance(balance int, transaction Transaction) int {
	if transaction.category == "credit" {
		return balance + transaction.amount
	}

	return balance - transaction.amount
}

func main() {
	var err error
	fmt.Println("Balance of account over time:\n")
	balance := 1000
	getNext := account.runningBalanceCalculator(balance, 0)

	for {
		fmt.Println(balance)
		balance, err = getNext()
		if err != nil {
			break
		}
	}

	fmt.Println("\nFinal balance:", balance)
}
