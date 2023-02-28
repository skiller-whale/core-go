package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Circumference() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Circumference() float64 {
	return 2 * c.Radius * math.Pi
}

type RightAngledTriangle struct {
	Width, Height float64
}

func (t RightAngledTriangle) Area() float64 {
	return t.Width * t.Height / 2
}

func (t RightAngledTriangle) Circumference() float64 {
	return t.Width + t.Height + math.Sqrt(t.Width*t.Width+t.Height*t.Height)
}

type Line struct {
	Length float64
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 5},
		RightAngledTriangle{Width: 10, Height: 5},
		Line{Length: 10},
	}

	for _, shape := range shapes {
		fmt.Printf("%T[%v] has aread of %f and circumference of %f\n", shape, shape, shape.Area(), shape.Circumference())
	}
}
