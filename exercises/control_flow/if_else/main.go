package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This seed is needed for 'random' results to change each time
	rand.Seed(time.Now().UnixNano())

	// `num` is a random number between 0 and 100
	var num int = rand.Intn(100)

	// TODO: Add your if clause here.

	fmt.Println("Your number is:", num)
}
