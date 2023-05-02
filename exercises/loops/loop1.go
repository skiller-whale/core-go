package main

import "fmt"

func main() {
	multiplesOf := 5

	// Printf will output on the same line until a line break is added
	fmt.Printf("Multiples of %v: ", multiplesOf)

	// TODO: Output multiples of 1-10 with a for loop
	fmt.Printf("%v ", 1*multiplesOf)
	fmt.Printf("%v ", 2*multiplesOf)

	// Called once to add a line break
	fmt.Println()

}
