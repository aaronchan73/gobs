package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/aaronchan73/gobs"
)

var counter *gobs.Counter = gobs.CreateCounter(0)

func main() {
	var wg sync.WaitGroup

	os.Setenv("COLLECTOR_ADDRESS", "http://localhost:8080")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			counter.IncrementCounter()
			counter.PrintCounter()
		}()
	}

	wg.Wait()
}
