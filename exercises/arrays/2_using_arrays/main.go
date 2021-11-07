package main

import (
  "fmt"
)

func main() {
	chunkedMessage := [6][6]byte{
		{34, 84, 105, 109, 101, 32},
		{102, 111, 114, 32, 99, 97},
		{107, 101, 33, 34, 44, 32},
		{114, 101, 112, 108, 105, 101},
		{0, 0, 0, 0, 0, 0},
		{100, 32, 83, 101, 98, 46},
	}
}

func printMessage(message [30]byte) {
  fmt.Println(string(message[:]))
}
