package main

import "fmt"

func main() {
	var badIndex int = 3

	// Create an array of length 2
	var shortArray = [2]int{1, 2}

	fmt.Println("The fourth item is:", shortArray[badIndex])
}
