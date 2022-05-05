package main

import (
	"fmt"
	"sync"
	"time"
)

type countTracker struct {
	counter int
}

func (t *countTracker) countUp(guarded bool) {
	if guarded {
		// Make sure only one goroutine at a time can execute the next line of code
	}
	t.counter += 1
}

func countToAMillion(guarded bool) int {
	if guarded {
		// Delete warning when finished
		fmt.Println("No guard implemented yet, don't trust the results!")
	}
	wg := sync.WaitGroup{}
	count := countTracker{}

	for i := 0; i < 1000000; i += 1 {
		wg.Add(1)
		go func() {
			count.countUp(guarded)
			wg.Done()
		}()
	}
	wg.Wait()

	return count.counter
}

func main() {
	start := time.Now()
	//count1 := countToAMillion(true)
	time1 := time.Since(start)
	count2 := countToAMillion(false)
	time2 := time.Since(start) - time1

	//fmt.Println("Guarded count:  ", count1, "in", time1)
	fmt.Println("Unguarded count:", count2, "in", time2)
}
