package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	gobs "github.com/aaronchan73/gobs/lib"
)

// Collector is the ingestion of data
type Collector struct {
	logs       map[int64][]gobs.Log
	counters   map[int64]gobs.Counter
	gauges     map[int64]gobs.Gauge
	histograms map[int64]gobs.Histogram
	traces     map[int64]gobs.Trace
	lock       *sync.RWMutex
}

const EXPORT_INTERVAL = 10

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
	collector.logs[log.ID] = append(collector.logs[log.ID], log)
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

// runExport exports collector data every EXPORT_INTERVAL seconds
func runExport(collector *Collector) {
	for {
		collector.lock.RLock()
		for _, logs := range collector.logs {
			for _, log := range logs {
				gobs.PrintLog(log)
			}
		}

		for _, counter := range collector.counters {
			counter.PrintCounter()
		}

		for _, gauge := range collector.gauges {
			gauge.PrintGauge()
		}

		for _, histogram := range collector.histograms {
			histogram.PrintHistogram()
		}

		for _, trace := range collector.traces {
			trace.PrintTrace()
		}

		data := map[string]interface{}{
			"logs":       collector.logs,
			"counters":   collector.counters,
			"gauges":     collector.gauges,
			"histograms": collector.histograms,
			"traces":     collector.traces,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		StoreJSONData(jsonData)

		collector.lock.RUnlock()

		time.Sleep(EXPORT_INTERVAL * time.Second)
	}
}

// exportData is the API handler for JSON data
func exportData(w http.ResponseWriter, r *http.Request) {
	data := GetJSONData()

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// StoreJSONData stores collector data into a JSON
func StoreJSONData(data []byte) {
	err := os.WriteFile("data.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

// GetJSONData retrieves JSON data from a file
func GetJSONData() []byte {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	collector = Collector{
		make(map[int64][]gobs.Log),
		make(map[int64]gobs.Counter),
		make(map[int64]gobs.Gauge),
		make(map[int64]gobs.Histogram),
		make(map[int64]gobs.Trace),
		&sync.RWMutex{},
	}

	go runExport(&collector)

	http.HandleFunc("/logs", updateLogs)
	http.HandleFunc("/counters", updateCounters)
	http.HandleFunc("/gauges", updateGauges)
	http.HandleFunc("/histograms", updateHistograms)
	http.HandleFunc("/traces", updateTraces)
	http.HandleFunc("/export", exportData)

	http.ListenAndServe(":8080", nil)
}
