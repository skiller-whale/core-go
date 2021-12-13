package main

import "fmt"

type Account struct {
	balance int
}

type Invoice struct {
	account Account
	amount  int
	ref     string
	settled bool
}

func main() {
	account := fetchDefault()

	fmt.Println("Initial balance:    ", account.balance)
	fmt.Println()

	invoices := []Invoice{
		Invoice{ref: "1", amount: 120, account: account},
		Invoice{ref: "2", amount: 220, account: account},
		Invoice{ref: "3", amount: 300, account: account},
		Invoice{ref: "4", amount: 200, account: account},
		Invoice{ref: "5", amount: 160, account: account},
	}

	for _, inv := range invoices {
		settleInvoice(inv)
		notify(inv)
	}
}

func settleInvoice(invoice Invoice) {
	newBalance := invoice.account.balance - invoice.amount
	if newBalance > 0 {
		invoice.settled = true
		invoice.account.balance = newBalance
	}
}

func notify(inv Invoice) {
	if inv.settled {
		fmt.Println("Settled", inv.ref, ":    ", inv.amount)
	} else {
		fmt.Println("Cannot settle invoice, insufficient balance.")
	}

	fmt.Println("    New balance:    ", inv.account.balance)
}

func fetchDefault() Account {
	return Account{balance: 500}
}
