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

// Gauge is a single numerical value
type Gauge struct {
	value int64
	lock  *sync.RWMutex
}

// createCounter creates a Counter
func createCounter() *Counter {
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

// createGauge creates a Gauge
func createGauge() *Gauge {
	gauge := Gauge{
		0,
		&sync.RWMutex{},
	}

	return &gauge
}

// updateGauge updates an existing Counter
func (gauge *Gauge) updateGauge(value int64) {
	gauge.lock.Lock()
	defer gauge.lock.Unlock()

	gauge.value = value
}

// printGauge prints an existing Gauge
func (gauge *Gauge) printGauge() {
	gauge.lock.RLock()
	defer gauge.lock.RUnlock()

	fmt.Println(gauge.value)
}
