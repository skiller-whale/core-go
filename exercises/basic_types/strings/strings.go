package main

import "fmt"

func main() {
	str := "café"

	fmt.Println([]rune(str))
	fmt.Println([]byte(str))

	fmt.Println("\nIterating over runes")
	for pos, char := range str {
		fmt.Println(pos, ":", string(char), char)
	}

	fmt.Println("\nIterating over bytes")
	for i := 0; i < len(str); i++ {
		fmt.Println(i, ":", string(str[i]), str[i])
	}
}
