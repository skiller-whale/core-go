package main

import (
	"fmt"
	"os"
)

type ArithmeticOp func(int, int) int

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Syntax: %s <op> <a> <b>\n", os.Args[0])
		return
	}

	var a, b int
	fmt.Sscan(os.Args[2], &a)
	fmt.Sscan(os.Args[3], &b)

	ops := map[string]Op{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { return a / b },
	}
	op, ok := ops[os.Args[1]]
	if ok {
		fmt.Println(op(a, b))
	} else {
		fmt.Println("Unknown op!")
	}
}
