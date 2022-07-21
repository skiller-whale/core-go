package main

import "fmt"

func main() {
	s := []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}

	for p := 2; p < 5; p++ {
		for q := p + p; q < 20; q += p {
			s[q] = true
		}
		for p++; s[p]; p++ {
		}
	}

	for i, v := range s {
		if !v {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
