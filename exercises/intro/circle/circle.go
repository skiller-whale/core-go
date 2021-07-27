package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int = 10
	var pi float64 = math.Pi
	var areaOfCircle float64 = pi * radius * radius
	fmt.Println("Area of a circle with radius 10 is:", areaOfCircle)
}
