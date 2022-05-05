package main

import "sync"
import "fmt"

func spawnNamePrinter(wg sync.WaitGroup, name string) {
	go func() {
		fmt.Println(name)
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup

	names := []string{"Ada", "Briony", "Cathy"}

	wg.Add(len(names))

	for _, name := range names {
		spawnNamePrinter(wg, name)
	}

	wg.Wait()
}
