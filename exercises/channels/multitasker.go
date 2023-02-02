// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func theNumbers(c chan []int) {
	defer close(c)
	list := []int{4, 8, 15, 16, 23, 42}
	for i := range list {
		c <- list[:i+1]
		time.Sleep(1500 * time.Millisecond)
	}
}

func progressBar(c chan float64) {
	defer close(c)
	for p := 0.0; p < 1.0; {
		time.Sleep(345 * time.Millisecond)
		p += rand.Float64() / 10.0
		if p > 1.0 {
			p = 1.0
		}
		c <- p
	}
}

func ticker(c chan time.Time) {
	for {
		time.Sleep(time.Second)
		c <- time.Now()
	}
}

// Like fmt.PrintLn, but moves cursor to start of given line#, then moves it back again after
func PrintlnOn(l int, a ...any) {
	esc := "\u001B["
	fmt.Print(esc + strconv.Itoa(1) + "G" + esc + strconv.Itoa(l) + "B")
	fmt.Println(a...)
	fmt.Print(esc + strconv.Itoa(l+1) + "A")
}

func main() {
	c1 := make(chan []int)
	c2 := make(chan float64)
	c3 := make(chan time.Time)

	go theNumbers(c1)
	go progressBar(c2)
	go ticker(c3)

	// TODO: write a loop that prints the progress of all 3 activities,
	// one per line using PrintlnOn.
	//
	// Quit when theNumbers and progressBar have both completed and closed
	// their channels.
}
