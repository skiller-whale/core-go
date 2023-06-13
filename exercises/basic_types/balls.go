package main

import "fmt"

func main() {
	var footballPitchLengthM, footballPitchWidthM uint8 = 105, 68 // in meters
	var golfballDiameterMm uint8 = 44                             // in millimeters, oops!

	// Only change this line, nothing above or below
	var golfballsOnFootballPitch int32 = footballPitchLengthM / golfballDiameterMm * footballPitchWidthM / golfballDiameterMm

	fmt.Println("You can fit", golfballsOnFootballPitch, "golf balls on a football pitch")
}
