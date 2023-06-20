package main

import (
	"fmt"
)

func PrintlnBig(args ...interface{}) {
	for _, arg := range args {
		// arg can be *anything*, handle with care
	}
	// use fmt.Print above, then add final newline
	fmt.Println()
}

func main() {
	PrintlnBig(99, "red balloons", "floating in the summer sky")
}
