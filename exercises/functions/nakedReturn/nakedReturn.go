package main

import "fmt"

func calculateRunningAverage() (int, int, error) {
	var (
		avg, count int
		err        error
	)
	for err == nil {
		var num int
		// read user input
		_, err = fmt.Scan(&num)
		// if user input is invalid
		if err != nil {
			return avg, count, err
		}
		avg = ((avg * count) + num) / (count + 1)
		count++
		fmt.Println("Average:", avg)
		fmt.Printf("Next number:")
	}
	return avg, count, err
}

func main() {
	fmt.Println("Input a number to calculate rolling average:")
	avg, count, err := calculateRunningAverage()
	fmt.Println("Average:", avg)
	fmt.Println("Total Numbers:", count)
	fmt.Println("Error:", err)
}
