package main

import "sync"
import "fmt"

func main() {
	var wg sync.WaitGroup

	names := []string{"Ada", "Briony", "Cathy"}

	wg.Add(len(names))

	for _, name := range names {
		go func() {
			fmt.Println(name)
			wg.Done()
		}()
	}

	wg.Wait()
}
