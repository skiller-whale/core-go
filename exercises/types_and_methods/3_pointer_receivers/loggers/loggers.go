package loggers

import (
	"fmt"
	"os"
)

type Message struct {
	Level string
	Body  string
}

//---------------------------------
// DualLogger and methods
//---------------------------------

type DualLogger struct {
	FileLogger
}

func (dl DualLogger) Log(msg Message) {
	info := dl.format(msg)
	dl.writeToFile(info)
	fmt.Println(info)
}

//---------------------------------
// FileLogger and methods
//---------------------------------

type FileLogger struct {
	Path string
	file *os.File
}

func (fl FileLogger) Log(msg Message) {
	info := fl.format(msg)
	fl.writeToFile(info)
}

func (fl FileLogger) format(msg Message) string {
	return fmt.Sprintf("[%v] %v", msg.Level, msg.Body)
}

// writeToFile sets the file field of a FileLogger to an open file if it is currently nil.
func (fl FileLogger) writeToFile(message string) {
	var err error

	if fl.file == nil {
		fl.file, err = os.OpenFile(fl.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		fmt.Println("File opened")
	}

	if err != nil {
		panic(err)
	}

	_, err = fl.file.WriteString(message + "\n")
	if err != nil {
		panic(err)
	}
}

func (fl FileLogger) Close() {
	if fl.file == nil {
		panic("No file to close!")
	}

	fl.file.Close()
}
