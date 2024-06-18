package main

import (
	"fmt"

	"github.com/skiller-whale/core-go/exercises/types_and_methods/3_pointer_receivers/loggers"
)

func main() {
	fl := loggers.FileLogger{Path: "production.log"}
	logger := loggers.DualLogger{FileLogger: fl}

	messages := []loggers.Message{
		loggers.Message{Level: "info", Body: "Process started"},
		loggers.Message{Level: "warn", Body: "Just a warning"},
		loggers.Message{Level: "info", Body: "Process ending"},
	}

	for _, msg := range messages {
		logger.Log(msg)
	}

	fmt.Println("Finished")
	logger.Close()
}
