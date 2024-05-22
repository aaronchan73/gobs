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

// Histogram is a range of values
type Histogram struct {
	buckets map[int64]int64
	lock    *sync.RWMutex
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

// createHistogram creates a Histogram
func createHistogram() *Histogram {
	histogram := Histogram{
		make(map[int64]int64),
		&sync.RWMutex{},
	}

	return &histogram
}

// updateHistogram updates an existing Histogram
func (histogram *Histogram) updateHistogram(value int64) {
	histogram.lock.Lock()
	defer histogram.lock.Unlock()

	if count, ok := histogram.buckets[value]; ok {
		histogram.buckets[value] = count + 1
	} else {
		histogram.buckets[value] = 1
	}
}

// printHistogram prints an existing Histogram
func (histogram *Histogram) printHistogram() {
	histogram.lock.RLock()
	defer histogram.lock.RUnlock()

	for value, count := range histogram.buckets {
		bucket := fmt.Sprintf("%d: %d\n", value, count)
		fmt.Println(bucket)
	}
}
