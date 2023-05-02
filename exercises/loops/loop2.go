package main

import "fmt"

func main() {
	multiplesOf := 5

	// Printf will output on the same line until a line break is added
	fmt.Printf("Multiples of %v: ", multiplesOf)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", i*multiplesOf)
	}

	// Called once to add a line break
	fmt.Println()

}
