package main

import (
	"fmt"
	"sort"
)

func main() {
	a := account

	dateComparator := func(i, j int) bool {
		return a.transactions[i].date.Before(a.transactions[j].date)
	}

	// TODO: Define the comparators here

	var action int

	getInput(&action)

	switch action {
	case 1:
		sort.Slice(a.transactions, dateComparator)
	case 2:
		sort.Slice(a.transactions, amountComparator)
	case 3:
		sort.Slice(a.transactions, categoryComparator)
	}
	printTransactions(a.transactions)
}

func getInput(action *int) {
	fmt.Println("Choose sorting method:")
	fmt.Println("1 - Date")
	fmt.Println("2 - Amount")
	fmt.Println("3 - Category")
	fmt.Scanln(action)
}
