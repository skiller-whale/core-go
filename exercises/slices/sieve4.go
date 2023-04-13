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

func summariseSieve(s []bool) (p []int) {
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
