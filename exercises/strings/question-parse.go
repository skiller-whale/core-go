package main

import "fmt"
import "io/ioutil"

//import "strings"
//import "strconv"

func readFileAsString(filename string) string {
	if content, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		return string(content)
	}
}

func main() {
	questions := readFileAsString("JeopardyData.tsv")
	fmt.Printf("Read %dKB\n", len(questions)/1024)
}
