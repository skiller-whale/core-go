package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func estimateOfPrimesIn(n int) int {
	return int(float64(n) / math.Log(float64(n)))
}

func sieve(n int) (s []bool) {
	s = make([]bool, n)
	s[0] = true

	for p := 2; p < int(math.Sqrt(float64(n))); {
		for q := p + p; q < n; q += p {
			s[q] = true
		}
		for p++; s[p]; p++ {
		}
	}

	return s
}

func summariseSieve(s []bool) (p []int) {
	p = make([]int, 0, estimateOfPrimesIn(len(s)))
	for i, notPrime := range s {
		if !notPrime {
			p = append(p, i)
		}
	}
	return p
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	mySieve := sieve(n)
	start := time.Now()
	summary := summariseSieve(mySieve)
	fmt.Println("There were", len(summary), "primes (summarised in", time.Since(start).Seconds(), "s)")
}
