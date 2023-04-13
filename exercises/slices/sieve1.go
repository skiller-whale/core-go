package main

func main() {
	// TODO: declare `notPrime`` as a "sieve" of 20 boolean values, defaulting to false

	for p := 2; p < len(notPrime); p++ {
		if !notPrime[p] {
			for q := 2; q*p < len(notPrime); q += 1 {
				// TODO: mark the number at q*p to as "notPrime" (not prime)
			}
		}
	}

	// TODO: print out the prime numbers we've found (where the value is false)
}
