package main

import (
	"fmt"
	"time"
)

type Log struct {
	path      string
	createdOn time.Time
	updatedOn time.Time
}

type GuessingGameClient struct {
	eventLog *Log
	score    int
}

func start(c GuessingGameClient) {
	fmt.Println("Starting Game")
	for i := 0; i < 5; i++ {
		randomNumber := zeroOrOne()
		guess, err := userInput()
		fmt.Println("Expected:", randomNumber, " Got:", guess)

		// Complete the function implementation
		if err != nil {
			event := "error"
		}
	}
}

func formatEventMessage(e string) string {
	return fmt.Sprintf("%v: %v\n", time.Now(), e)
}

func (l *Log) append(event string) {
	eventMsg := formatEventMessage(event)
	f := openFile(l.path)
	defer closeFile(f)
	writeToFile(f, eventMsg)
	l.updatedOn = time.Now()
}

func main() {
	c := newClient()
	start(c)
	fmt.Println("Final score:", c.score)
}
