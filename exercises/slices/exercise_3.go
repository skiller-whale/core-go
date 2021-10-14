package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/slices/names"
)

func main() {
	popularNames := names.LoadNames()

	var mostPopular []string

	// append the first five names from `popularNames` to `mostPopular` here

	fmt.Println("\t", mostPopular)
	fmt.Println("\tmostPopular capacity: ", "...")

	fmt.Printf("\tpopularNames capacity: %v, ", "...")
	fmt.Printf("length: %v", "...")
}
