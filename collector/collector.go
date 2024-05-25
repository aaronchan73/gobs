package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/aaronchan73/gobs"
)

// Collector is the ingestion of data
type Collector struct {
	logs       map[int64]gobs.Log
	counters   map[int64]gobs.Counter
	gauges     map[int64]gobs.Gauge
	histograms map[int64]gobs.Histogram
	traces     map[int64]gobs.Trace
	lock       *sync.RWMutex
}

var collector Collector

// updateLogs receives logs via HTTP and updates collector
func updateLogs(w http.ResponseWriter, r *http.Request) {
	collector.lock.Lock()
	defer collector.lock.Unlock()

	var log gobs.Log

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &log)
	gobs.PrintLog(log)
	collector.logs[log.ID] = log
}

// updateCounters receives counters via HTTP and updates collector
func updateCounters(w http.ResponseWriter, r *http.Request) {
	collector.lock.Lock()
	defer collector.lock.Unlock()

	var counter gobs.Counter

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &counter)
	counter.PrintCounter()
	collector.counters[counter.ID] = counter
}

// updateGauges receives gauges via HTTP and updates collector
func updateGauges(w http.ResponseWriter, r *http.Request) {
	collector.lock.Lock()
	defer collector.lock.Unlock()

	var gauge gobs.Gauge

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &gauge)
	gauge.PrintGauge()
	collector.gauges[gauge.ID] = gauge
}

// updateHistograms receives histograms via HTTP and updates collector
func updateHistograms(w http.ResponseWriter, r *http.Request) {
	collector.lock.Lock()
	defer collector.lock.Unlock()

	var histogram gobs.Histogram

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &histogram)
	histogram.PrintHistogram()
	collector.histograms[histogram.ID] = histogram
}

// updateTraces receives traces via HTTP and updates collector
func updateTraces(w http.ResponseWriter, r *http.Request) {
	collector.lock.Lock()
	defer collector.lock.Unlock()

	var trace gobs.Trace

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &trace)
	trace.PrintTrace()
	collector.traces[trace.ID] = trace
}

func main() {
	collector = Collector{
		make(map[int64]gobs.Log),
		make(map[int64]gobs.Counter),
		make(map[int64]gobs.Gauge),
		make(map[int64]gobs.Histogram),
		make(map[int64]gobs.Trace),
		&sync.RWMutex{},
	}

	http.HandleFunc("/logs", updateLogs)
	http.HandleFunc("/counters", updateCounters)
	http.HandleFunc("/gauges", updateGauges)
	http.HandleFunc("/histograms", updateHistograms)
	http.HandleFunc("/traces", updateTraces)

	http.ListenAndServe(":8080", nil)
}
