package main

import "fmt"

func main() {
	var numberOfMinutes uint8 = 200
	var numberOfSecInAMin int8 = 60
	totalNumberOfSeconds := numberOfMinutes * numberOfSecInAMin
	fmt.Println("Total number of seconds", totalNumberOfSeconds)

	var numberOfSecInAnHour int8 = 3600
	totalNumberOfHours := totalNumberOfSeconds / numberOfSecInAnHour
	fmt.Println("Total number of hours", totalNumberOfHours)
}
