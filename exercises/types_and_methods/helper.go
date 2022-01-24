package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func newClient() GuessingGameClient {
	return GuessingGameClient{eventLog: newLog()}
}

func otherClient() GuessingGameClient {
	return GuessingGameClient{}
}

func newLog() *Log {
	f, err := createFile("event.log")
	if err != nil {
		panic(err)
	}
	defer closeFile(f)
	return &Log{path: "event.log", createdOn: time.Now()}
}

func createFile(path string) (*os.File, error) {
	return os.Create(path)
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func writeToFile(f *os.File, data string) {
	_, err := f.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func closeFile(f *os.File) {
	f.Close()
}

func getFileContent(path string, n int) []string {
	n++
	f := openFile(path)
	defer closeFile(f)
	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	if len(lines) < n {
		return lines
	}
	return lines[len(lines)-n:]
}

func zeroOrOne() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}

func userInput() (int, error) {
	var n int
	fmt.Println("\nEnter your guess:\n")
	_, err := fmt.Scanln(&n)
	if err == nil && (n < 0 || n > 10) {
		err = errors.New("Out of range")
	}
	return n, err
}
