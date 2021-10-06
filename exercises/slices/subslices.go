package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	popularNames := loadNames()

	fmt.Println("The [name,count] for the 4th to 6th (inclusive) most popular were:", "...")
	fmt.Println("The [name,count] for the 3 least popular were:", "...")
}

// This is a utility function that loads the popular names from a CSV file.
// You do not need to edit anything in this function and are not expected
// to understand how it works.
func loadNames() [][]string {
	file, err := os.Open("popular_female_names_1990s.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records[1:]
}
