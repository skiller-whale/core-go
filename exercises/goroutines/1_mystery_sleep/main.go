package main

import "time"
import "fmt"

func main() {
	for n := 0; n < 10; n += 1 {
		go fmt.Print(n)
	}
	time.Sleep(time.Millisecond)
}
