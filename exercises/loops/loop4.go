package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Fragment 1: Read some input from the keyboard as a string
	var input string
	fmt.Println("Input: ")
	fmt.Scan(&input) // read something from the keyboard

	// Fragment 2: Convert a string into an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		// In this case, number is 0
	}

	//
	// Put these fragments 👆🏻 inside the loop below
	//
	for multiplesOf := 1; multiplesOf <= 10; multiplesOf++ {
		//
		// Instead of showing the multiples of 1-10 each time, ask the user what number
		// they would like to see multiples of.
		//
		// If they enter "quit", exit the program
		// If they enter an invalid number, ignore their input
		// If they enter a valid number, show its multiples from 1-10
		//
		fmt.Printf("Multiples of %v: ", multiplesOf)
		for i := 1; i <= 10; i++ {
			fmt.Printf("%v ", i*multiplesOf)
		}
		fmt.Println()
	}
}
