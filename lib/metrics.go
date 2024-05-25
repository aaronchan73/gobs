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

// CreateCounter creates a Counter
func CreateCounter() *Counter {
	counter := Counter{
		0,
		&sync.RWMutex{},
	}

	return &counter
}

// IncrementCounter increments an existing Counter
func (counter *Counter) IncrementCounter() {
	counter.lock.Lock()
	defer counter.lock.Unlock()

	counter.count++
}

// PrintCounter prints an existing Counter
func (counter *Counter) PrintCounter() {
	counter.lock.RLock()
	defer counter.lock.RUnlock()

	fmt.Println(counter.count)
}

// CreateGauge creates a Gauge
func CreateGauge() *Gauge {
	gauge := Gauge{
		0,
		&sync.RWMutex{},
	}

	return &gauge
}

// UpdateGauge updates an existing Counter
func (gauge *Gauge) UpdateGauge(value int64) {
	gauge.lock.Lock()
	defer gauge.lock.Unlock()

	gauge.value = value
}

// PrintGauge prints an existing Gauge
func (gauge *Gauge) PrintGauge() {
	gauge.lock.RLock()
	defer gauge.lock.RUnlock()

	fmt.Println(gauge.value)
}

// CreateHistogram creates a Histogram
func CreateHistogram() *Histogram {
	histogram := Histogram{
		make(map[int64]int64),
		&sync.RWMutex{},
	}

	return &histogram
}

// UpdateHistogram updates an existing Histogram
func (histogram *Histogram) UpdateHistogram(value int64) {
	histogram.lock.Lock()
	defer histogram.lock.Unlock()

	if count, ok := histogram.buckets[value]; ok {
		histogram.buckets[value] = count + 1
	} else {
		histogram.buckets[value] = 1
	}
}

// PrintHistogram prints an existing Histogram
func (histogram *Histogram) PrintHistogram() {
	histogram.lock.RLock()
	defer histogram.lock.RUnlock()

	for value, count := range histogram.buckets {
		bucket := fmt.Sprintf("%d: %d\n", value, count)
		fmt.Println(bucket)
	}
}
