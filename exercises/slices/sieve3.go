package main

import (
	"fmt"
	"math"
)

func sieve(n int) []bool {
	s := make([]bool, n)
	s[0] = true

	for p := 2; p < int(math.Sqrt(float64(n))); p++ {
		for q := p + p; q < n; q += p {
			s[q] = true
		}
		for p++; s[p]; p++ {
		}
	}

	return s
}

func main() {
	for i, v := range sieve(100) {
		if !v {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
