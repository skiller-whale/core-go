package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	date     time.Time
	amount   int
	category string
}

type Account struct {
	transactions []Transaction
}

var account = Account{
	transactions: []Transaction{
		{
			date: time.Date(2022, time.February, 22, 22, 22, 22, 22, time.UTC),
			category: "credit",
			amount:   210,
		},
		{
			date: time.Date(2022, time.February, 25, 20, 12, 2, 32, time.UTC),
			category: "debit",
			amount:   100,
		},
		{
			date: time.Date(2022, time.February, 26, 10, 2, 3, 3, time.UTC),
			category: "debit",
			amount:   1000,
		},
		{
			date: time.Date(2022, time.February, 27, 9, 6, 3, 3, time.UTC),
			category: "credit",
			amount:   560,
		},
		{
			date: time.Date(2022, time.February, 28, 8, 14, 9, 14, time.UTC),
			category: "debit",
			amount:   45,
		},
	},
}

func printSummary(t Transaction) {
	fmt.Printf("%v was %ved on %v\n", t.amount, t.category, t.date.Format(time.RFC1123))
}

func printTransactions(transactions []Transaction) {
	for _, t := range transactions {
		printSummary(t)
	}
}
