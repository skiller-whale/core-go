package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/slices/names"
)

func main() {
	popularNames := names.LoadNames()

	fmt.Println("The [name,count] for the 4th to 6th (inclusive) most popular were:", "...")
	fmt.Println("The [name,count] for the 3 least popular were:", "...")
}
