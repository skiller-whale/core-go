package main

import "fmt"

func main() {

	for multiplesOf := 1; multiplesOf <= 10; multiplesOf++ {
		fmt.Printf("Multiples of %v: ", multiplesOf)
		for i := 1; i <= 10; i++ {
			fmt.Printf("%v ", i*multiplesOf)
		}
		fmt.Println()
	}
}
