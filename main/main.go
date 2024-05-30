package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	gobs "github.com/aaronchan73/gobs/lib"
)

const CLIENT_ID = 0

var counter *gobs.Counter = gobs.CreateCounter(CLIENT_ID)
var gauge *gobs.Gauge = gobs.CreateGauge(CLIENT_ID)
var histogram *gobs.Histogram = gobs.CreateHistogram(CLIENT_ID)
var trace *gobs.Trace = gobs.CreateTrace(CLIENT_ID)

func main() {
	var wg sync.WaitGroup

	os.Setenv("COLLECTOR_ADDRESS", "http://collector-service.default.svc.cluster.local:8080")

	log := gobs.CreateLog(0, "Running main function")
	gobs.UpdateLog(log)
	gobs.PrintLog(log)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		log := gobs.CreateLog(0, fmt.Sprintf("Running iteration: %d", i))
		gobs.UpdateLog(log)
		gobs.PrintLog(log)

		go func(increment int64) {
			defer wg.Done()

			start := time.Now()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			end := time.Since(start)

			counter.IncrementCounter()
			gauge.UpdateGauge(increment)
			histogram.UpdateHistogram(increment)
			span := gobs.CreateSpan(increment, "main", end)
			trace.UpdateTrace(span)
		}(int64(i))
	}

	wg.Wait()

	log = gobs.CreateLog(0, "Completed main function")
	gobs.UpdateLog(log)
	gobs.PrintLog(log)

	counter.PrintCounter()
	gauge.PrintGauge()
	histogram.PrintHistogram()
	trace.PrintTrace()
}
