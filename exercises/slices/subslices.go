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

// To be de-duplicated by pulling out to separate file / package
func loadNames() [][]string {
	file, err := os.Open("popular_female_names_1990s.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records[1:]
}
