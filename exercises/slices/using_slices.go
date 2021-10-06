package main

// import the names data
import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// First, uncomment the line below
	//popularNames := loadNames()

	fmt.Println()
	fmt.Println("\tTotal number of girls names:", "...")
	fmt.Println("\tThe most popular girls name was", "...")
	fmt.Println("\tThe fourth most popular girls name was", "...")
	fmt.Println("\tThe number of times that the second least popular girls name was chosen was", "...")

	fmt.Println()
	fmt.Println("(End of Using Slices exercise)")
	fmt.Println()

	var mostPopular []string

	// append the first five names from `popularNames` to `mostPopular` here

	fmt.Println("\t", mostPopular)
	fmt.Println("\tmostPopular capacity: ", "...")

	fmt.Printf("\tpopularNames capacity: %v, ", "...")
	fmt.Printf("length: %v", "...")
	fmt.Println()
}

// To be pulled out to separate file / package
func loadNames() [][]string {

	file, err := os.Open("popular_female_names_1990s.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records[1:]
}
