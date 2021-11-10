package main

import (
	"fmt"
)

func main() {
	message := [31]byte{
		34, 65, 110, 100, 32, 97,
		32, 99, 117, 112, 32, 111,
		102, 32, 116, 101, 97, 63,
		34, 44, 32, 90, 111, 101,
		32, 97, 115, 107, 101, 100,
		46,
	}

	longerName := [5]byte{66, 101, 108, 108, 97}

	fmt.Println(string(message[:]))
	fmt.Println("New name:", string(longerName[:]))

	longerMessage := make([]byte, 33)

	// Populate longerMessage here!

  fmt.Println(string(longerMessage[:]))
}
