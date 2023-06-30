package main

import "fmt"
import "time"

// import "strings"

func main() {
	t0 := time.Now()
	// 1
	fmt.Println(time.Since(t0))

	t1 := time.Now()
	// 2
	fmt.Println(time.Since(t1))

	t2 := time.Now()
	// 3
	fmt.Println(time.Since(t2))
}
