package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/slices/names"
)

func main() {
	popularNames := names.LoadNames()

	fmt.Println("First 8 in popularNames:")
	names.PrintNames(popularNames[0:8])
}
