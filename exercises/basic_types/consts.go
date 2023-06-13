package main

import "fmt"

func main() {
	const pi = 3.14159265358979323846264338327950288419716939937510582097494459
	const radius = 8

	// 1) Why does this not compile?
	radiusSquared := radius * radius
	area := pi * radiusSquared

	// But this does?
	//
	// 2) area := pi * radius * radius

	// 3) Can you change the definition of `radius` to make the first version compile?

	fmt.Println("Area of circle of radius", radius, "is", area)
}
