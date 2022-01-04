package main

import "fmt"

func main() {
	str := "Whale, hello there"
	// Print memory address here

	exclaim(str)
	// Print memory address here

	fmt.Println(str)
}

func exclaim(str string) {
	str = str + exclamationMark
}

const exclamationMark = "!"
const quotationMark = "\""
