package main

import (
	"fmt"
	//  "strings"
)

type Message struct {
	Level string
	Body  string
}

func main() {
	messages := []Message{
		Message{Level: "debug", Body: "Setting up..."},
		Message{Level: "info", Body: "Process started"},
		Message{Level: "warn", Body: "Just a warning"},
		Message{Level: "info", Body: "Process ending"},
		Message{Level: "debug", Body: "Tidying up..."},
	}

	for _, m := range messages {
		formattedLevel := fmt.Sprintf("[%v]", m.Level)
		fmt.Printf("%v %v\n", formattedLevel, m.Body)
	}

	fmt.Println("Finished")
}
