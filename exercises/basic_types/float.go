package main

import "fmt"

func main() {
	var f1 float32 = 1.0
	var f2 float32 = 0.0
	var f3 = f1 / f2
	var f4 = f2 / f2
	fmt.Println(f3)
	fmt.Println(f4)

	var i1 int = 1
	var i2 int = 0
	var i3 = i1 / i2
	fmt.Println(i3)
}
