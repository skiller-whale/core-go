package main

import "fmt"

type Account struct {
	balance *int
}

func main() {
	minimum := balance(500)
	var accounts []*Account

	accounts = loadAll()

	for i, a := range accounts {
		increaseBalance(a.balance)
		if a.balance == minimum {
			fmt.Println("Account", i, "at minimum")
		} else {
			fmt.Println("Account", i, ":", *a.balance)
		}
	}
}

func increaseBalance(balance *int) {
	*balance += 500
}

func balance(b int) *int {
	return &b
}
