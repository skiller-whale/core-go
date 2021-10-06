package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	popularNames := loadNames()

	fmt.Println("First 8 in popularNames:")
	printNames(popularNames[0:8])
}

// The following are utility functions used in the exercise.
// You do not need to edit anything below and are not expected
// to understand how the functions work.
func loadNames() [][]string {

	file, err := os.Open("popular_female_names_1990s.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records[1:]
}

func printNames(names [][]string) {
	for _, name := range names {
		printName(name)
	}
	fmt.Println()
}

func printName(name []string) {
	fmt.Printf("%v\t", name[0])
	if len(name[0]) < 8 {
		fmt.Printf("\t")
	}
	fmt.Printf("(%v)", name[1])
	fmt.Println()
}
