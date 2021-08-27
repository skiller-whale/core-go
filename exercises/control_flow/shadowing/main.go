package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 10
	// It's much more common to see blocks with an if, switch or for statement.
	// This example uses plain {  } blocks for simplicity.
	{
		y := "twenty"
		{
			fmt.Println(x, y)
			x := "three"
			y = "thirty"
			fmt.Println(x, y)
		}
		fmt.Println(x, y)
	}
	fmt.Println(x, y)
}
