package main

import "fmt"

const Fizz = 1
const Buzz = 2

func check(n int) (o int) {
	if n%3 == 0 {
		o |= Fizz
	}
	if n%5 == 0 {
		o |= Buzz
	}
	return o
}

func main() {
	for i := 1; i < 26; i++ {
		o := check(i)
		fmt.Print(i)
		if o&Fizz != 0 {
			fmt.Print(" fizz")
		}
		if o&Buzz != 0 {
			fmt.Print(" buzz")
		}
		fmt.Println()
	}
}
