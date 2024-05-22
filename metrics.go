package gobs

import (
	"fmt"
	"sync"
)

// Counter is a monotonically increasing count
type Counter struct {
	count int64
	lock  *sync.RWMutex
}

// createCounter creates a Counter
func createCounter(message string) *Counter {
	counter := Counter{
		0,
		&sync.RWMutex{},
	}

	return &counter
}

// incrementCounter increments an existing Counter
func (counter *Counter) incrementCounter() {
	counter.lock.Lock()
	defer counter.lock.Unlock()

	counter.count++
}

// printCounter prints an existing Counter
func (counter *Counter) printCounter() {
	counter.lock.RLock()
	defer counter.lock.RUnlock()

	fmt.Println(counter.count)
}
