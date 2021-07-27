package main

import "fmt"

func main() {
	var badIndex int = 3
	var shortArray = [1]string{"Hello"}

	fmt.Println("The fourth item is:", shortArray[badIndex])
}
