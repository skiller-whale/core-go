package main

type Transaction struct {
	category string
	amount   int
}

type Account struct {
	transactions []Transaction
}

var account = Account{
	transactions: []Transaction{
		{
			category: "credit",
			amount:   210,
		},
		{
			category: "debit",
			amount:   100,
		},
		{
			category: "debit",
			amount:   1000,
		},
		{
			category: "credit",
			amount:   560,
		},
		{
			category: "debit",
			amount:   45,
		},
	},
}
