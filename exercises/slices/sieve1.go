package main

func main() {
	// TODO: declare s as a "sieve" of 20 boolean values, defaulting to false

	for p := 2; p < 5; p++ {
		for q := p + p; q < len(s); q += p {
			// TODO: set index q to true, indicating it is not prime
		}
		for p++; s[p]; p++ {
		}
	}

	// TODO: print out the prime numbers we've found (where the value is false)
}
