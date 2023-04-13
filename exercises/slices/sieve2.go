package main

import "fmt"

func main() {
	composite := []bool{true, true, 20: false}

	for p := 2; p < 21; p++ {
		if !notPrime[p] {
			for q := 2; q*p < 21; q += 1 {
				notPrime[q*p] = true
			}
		}
	}

	for i, v := range notPrime {
		if !v {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
