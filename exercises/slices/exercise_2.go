package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/slices/names"
)

func main() {
	popularNames := names.LoadNames()

	fmt.Println("\tTotal number of girls names:", "...")
	fmt.Println("\tThe most popular girls name was", "...")
	fmt.Println("\tThe fourth most popular girls name was", "...")
	fmt.Println(
		"\tThe second least popular girls name was",
		"...",
		"it was chosen",
		"...",
		"times",
	)
}
