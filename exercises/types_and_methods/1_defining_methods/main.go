package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/types_and_methods/1_defining_methods/loggers"
)

func main() {
	logger := loggers.FileLogger{Path: "production.log"}

	messages := []loggers.Message{
		loggers.Message{Level: "info", Body: "Process started"},
		loggers.Message{Level: "warn", Body: "Just a warning"},
		loggers.Message{Level: "info", Body: "Process ending"},
	}

	// Add range loop here

	fmt.Println("Finished")
}
