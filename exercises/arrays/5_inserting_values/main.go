
package main

import (
	"fmt"
)

func main() {
	chunkedMessage := [5][6]byte{
		{34, 84, 105, 109, 101, 32},
		{102, 111, 114, 32, 99, 97},
		{107, 101, 33, 34, 44, 32},
		{114, 101, 112, 108, 105, 101},
		{100, 32, 83, 101, 98, 46},
	}

	var message [30]byte

	for i, bytes := range chunkedMessage {
	        _ = copy(message[i*6:], bytes[:])
	}

	fmt.Println(string(message[:]))

	newName := [6]byte{74, 111, 115}

	// Replace the name, which needs to be a slice from newName
  c := copy(message[26:29], newName[:])

	fmt.Println(c)
	fmt.Println(string(message[:]))

	longerName := [5]byte{66, 101, 108, 108, 97}
	longerMessage := make([]byte, len(message) + 2)

	_ = copy(longerMessage[:], message[:26])
	_ = copy(longerMessage[26:], longerName[:])
	_ = copy(longerMessage[26 + len(longerName):], message[29:])
}
