package main

import "fmt"

func main() {
	var width, height, depth, litresInMCubed uint16 = 50, 25, 2, 1000
	fmt.Println("There is", width*height*depth*litresInMCubed, "ℓ of water in a regulation Olympic swimming pool")

	var quarterPounderWithCheeseCalories int16 = 518
	fmt.Println("Blue whales eat the equivalent of", 20000000/quarterPounderWithCheeseCalories, "quarter pounders with cheese per day")

	var annualTennisBallsAtWimbledon uint16 = 55000
	var annualTennisBallsAtUSOpen uint16 = 71001
	var annualTennisBallsAtRolandGarros uint16 = 65000
	var annualTennisBallsAtAustralianOpen uint16 = 48000
	fmt.Println("Average number of tennis balls used at a grand slam tennis tournament", (annualTennisBallsAtWimbledon+annualTennisBallsAtUSOpen+annualTennisBallsAtRolandGarros+annualTennisBallsAtAustralianOpen)/4)
}
