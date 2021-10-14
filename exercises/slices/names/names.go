// package names contains utility functions to load and print names from a csv
// file, which is included with the package.
package names

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"runtime"
)

func LoadNames() [][]string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Problem accessing names package")
		return [][]string{}
	}

	filePath := path.Join(path.Dir(thisFile), "popular_female_names_1990s.csv")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}

	return records[1:]
}

func PrintNames(names [][]string) {
	for _, name := range names {
		PrintName(name)
	}
	fmt.Println()
}

func PrintName(name []string) {
	fmt.Printf("%v\t", name[0])
	if len(name[0]) < 8 {
		fmt.Printf("\t")
	}
	fmt.Printf("(%v)", name[1])
	fmt.Println()
}
