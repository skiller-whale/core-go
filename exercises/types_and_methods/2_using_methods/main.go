package main

import (
	"github.com/skiller-whale/core-go/exercises/types_and_methods/2_using_methods/loggers"
)

func main() {
	fl := loggers.FileLogger{Path: "production.log"}

	messages := []loggers.Message{
		loggers.Message{Level: "info", Body: "Process started"},
		loggers.Message{Level: "warn", Body: "Just a warning"},
		loggers.Message{Level: "info", Body: "Process ending"},
	}

	logger := loggers.DualLogger{
		FileLogger: fl,
	}

	for _, msg := range messages {
		// Use DualLogger.Log here
		logger.Log(msg)
	}
}
