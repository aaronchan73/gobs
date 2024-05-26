package gobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Counter is a monotonically increasing count
type Counter struct {
	ID    int64 `json:"id"`
	Count int64 `json:"count"`
}

// Gauge is a single numerical value
type Gauge struct {
	ID    int64 `json:"id"`
	Value int64 `json:"value"`
}

// Histogram is a range of values
type Histogram struct {
	ID      int64           `json:"id"`
	Buckets map[int64]int64 `json:"buckets"`
}

// CreateCounter creates a Counter
func CreateCounter(id int64) *Counter {
	counter := Counter{
		id,
		0,
	}

	return &counter
}

// IncrementCounter increments an existing Counter
func (counter *Counter) IncrementCounter() {
	counter.Count++

	jsonBody, _ := json.Marshal(map[string]int64{
		"id":    counter.ID,
		"count": counter.Count,
	})
	responseBody := bytes.NewBuffer(jsonBody)

	requestURL := os.Getenv("COLLECTOR_ADDRESS") + "/counters"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}

// PrintCounter prints an existing Counter
func (counter *Counter) PrintCounter() {
	fmt.Println(counter.Count)
}

// CreateGauge creates a Gauge
func CreateGauge(id int64) *Gauge {
	gauge := Gauge{
		id,
		0,
	}

	return &gauge
}

// UpdateGauge updates an existing Counter
func (gauge *Gauge) UpdateGauge(value int64) {
	gauge.Value = value

	jsonBody, _ := json.Marshal(map[string]int64{
		"id":    gauge.ID,
		"value": gauge.Value,
	})
	responseBody := bytes.NewBuffer(jsonBody)

	requestURL := os.Getenv("COLLECTOR_ADDRESS") + "/gauges"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}

// PrintGauge prints an existing Gauge
func (gauge *Gauge) PrintGauge() {
	fmt.Println(gauge.Value)
}

// CreateHistogram creates a Histogram
func CreateHistogram(id int64) *Histogram {
	histogram := Histogram{
		id,
		make(map[int64]int64),
	}

	return &histogram
}

// UpdateHistogram updates an existing Histogram
func (histogram *Histogram) UpdateHistogram(value int64) {
	if count, ok := histogram.Buckets[value]; ok {
		histogram.Buckets[value] = count + 1
	} else {
		histogram.Buckets[value] = 1
	}

	jsonBody, _ := json.Marshal(map[string]interface{}{
		"id":      histogram.ID,
		"buckets": histogram.Buckets,
	})
	responseBody := bytes.NewBuffer(jsonBody)

	requestURL := os.Getenv("COLLECTOR_ADDRESS") + "/histograms"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}

// PrintHistogram prints an existing Histogram
func (histogram *Histogram) PrintHistogram() {
	for value, count := range histogram.Buckets {
		bucket := fmt.Sprintf("%d: %d", value, count)
		fmt.Println(bucket)
	}
}
