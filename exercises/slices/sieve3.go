package main

import (
	"fmt"
)

func sieve(n int) (notPrime []bool) {
	notPrime = make([]bool, n)
	notPrime[0] = true
	notPrime[1] = true

	for p := 2; p < n; p++ {
		if !notPrime[p] {
			for q := 2; q*p < n; q += 1 {
				notPrime[q*p] = true
			}
		}
	}

	return notPrime
}

func main() {
	for i, v := range sieve(100) {
		if !v {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
