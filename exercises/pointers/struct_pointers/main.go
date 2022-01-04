package main

import "fmt"

type Account struct {
	balance int
	ref   string
}

func main() {
	var account Account

	account = load()

	printAccount(account)
}

func load() Account {
	return Account{
		balance: 0,
		ref:   "ABC1234",
	}
}

func printAccount(account Account) {
	fmt.Println("Account:", account.balance, account.ref)
}
