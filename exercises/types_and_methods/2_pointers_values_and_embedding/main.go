package main

import (
	"github.com/skiller-whale/core-go/exercises/types_and_methods/2_pointers_values_and_embedding/loggers"
)

func main() {
	messages := []loggers.Message{
		loggers.Message{Level: "info", Body: "Process started"},
		loggers.Message{Level: "warn", Body: "Just a warning"},
		loggers.Message{Level: "info", Body: "Process ending"},
	}

	fl := loggers.FileLogger{Path: "production.log"}

	// Define DualLogger here

	for _, msg := range messages {
		// Use DualLogger.Log here
	}
}
