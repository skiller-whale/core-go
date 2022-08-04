package main

import "fmt"

// Uncomment from Exercise 2 onwards:
// import "github.com/skiller-whale/core-go/exercises/maps/countries"

// getFirstLetter returns the first character of a string.
// If the input string is empty, then the empty string will be returned.
func getFirstLetter(s string) string {
	r := []rune(s)
	return string(r[:1])
}

// printCountryName should print the name associated with a country code, or "Not found"
func printCountryName(codeToName map[string]string) {
	fmt.Print("Input a country code: ")
	var input string
	fmt.Scan(&input)

	// TODO: Implement the rest of this function
}

func main() {
	// Exercise 1: Create a map of codes to country names

	// Exercise 2: Add new entries to codeToName

	// Exercise 3: Find how many country codes are stored in codeToName

	// Exercise 4: User enters a code, print the respective country else print not found
	// printCountryName(c)

	// Exercise 5: Create a map of beginning letter to slice of codes

	//fmt.Println("Total number of country codes beginning with M:", "TODO")
	//fmt.Println("Total number of country codes beginning with Z:", "TODO")
}
