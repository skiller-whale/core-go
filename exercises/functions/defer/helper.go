package main

import (
	"errors"
	"fmt"
	"os"
)

var employees = []string{"2000", "1001", "1002", "1003"}
var opened = false

var employeeInfo = map[string]string{
	"1001": `{id: "1001", name: "Bruce Shrimpsteen"}`,
	"1002": `{id: "1002", name: "Eelton John"}`,
	"1003": `{id: "1003", name: "Tuna Turner"}`,
}

// Mock function that fetches info from api
func getInfo(employeeId string) (string, error) {
	if _, ok := employeeInfo[employeeId]; !ok {
		err := errors.New("employee not found")
		return "", err
	}
	return employeeInfo[employeeId], nil
}

// Mock function that uploads file to server
func uploadFile(file *os.File) error {
	return nil
}

func openFile(filePath string) *os.File {
	if opened {
		panic(filePath + " already open")
	}
	opened = true
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File, data string) (int, error) {
	fmt.Println("Writing to file:", data)
	return fmt.Fprintln(file, data)
}

func closeFile(file *os.File) {
	if !opened {
		panic("Cannot close unopened file")
	}
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	opened = false
}

func cleanFile(filePath string) {
	os.WriteFile(filePath, []byte{}, os.FileMode(0644))
}
