package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Syntax: %s <op> <a> <b>\n", os.Args[0])
		return
	}

	var a, b int
	fmt.Sscan(os.Args[2], &a)
	fmt.Sscan(os.Args[3], &b)

	switch os.Args[1] {
	case "add":
		fmt.Println(add(a, b))
	case "sub":
		fmt.Println(sub(a, b))
	case "mul":
		fmt.Println(mul(a, b))
	case "div":
		fmt.Println(div(a, b))
	default:
		fmt.Println("Unknown op!")
	}
}
