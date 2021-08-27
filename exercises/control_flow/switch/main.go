package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skiller-whale/core-go/exercises/control_flow/switch/dice"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	score := 0
	rounds := 5

	for i := 0; i < rounds; i++ {
		diceRoll := dice.RollDice() // a random integer between 1 and 6 (inclusive)
		guessed := dice.Guess()     // a user guess between 1 and 6

		// TODO - Switch statement goes here

		fmt.Println("The dice roll was", diceRoll)
		fmt.Println("You guessed:", guessed)
		fmt.Println("Total Score:", score)
		fmt.Println("\n---------------------------\n")
	}
}
