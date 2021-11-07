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


}

// Exercise 5 - Inserting Values
//
//   * At the end of the `main` function, declare a new byte array, `longName`,
//     containing the values: `66 101 108 108 97`
//
//   * Replace the three-byte name in `message` with this five byte name.
//     (The starting index of the name is `26`.)
//
