package loggers

import (
	"fmt"
	"os"
)

type Message struct {
	Level string
	Body  string
}

type FileLogger struct {
	Path         string
}

func (l FileLogger) writeToFile(message string) {
	file, err := os.OpenFile(l.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(message + "\n")
	if err != nil {
		panic(err)
	}
}

func (l FileLogger) format(msg Message) string {
	return fmt.Sprintf("[%v] %v", msg.Level, msg.Body)
}
