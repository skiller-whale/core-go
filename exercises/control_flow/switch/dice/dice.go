package dice

import (
	"fmt"
	"math/rand"
	"strconv"
)

func RollDice() int {
	return rand.Intn(6) + 1
}

func Guess() int {
	fmt.Println("Guess the number (1-6):")
	var input string
	for {
		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > 6 {
			fmt.Println("Please enter a number between 1 and 6")
			continue
		}
		return num
	}
}
