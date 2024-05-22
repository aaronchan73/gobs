package main

import (
	"gobs"
	"sync"
)

var counter *gobs.Counter = gobs.CreateCounter()

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncrementCounter()
			counter.PrintCounter()
		}()
	}

	wg.Wait()
}
