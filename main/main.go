package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/aaronchan73/gobs"
)

var counter *gobs.Counter = gobs.CreateCounter(0)

func main() {
	var wg sync.WaitGroup

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

	postBody, _ := json.Marshal(map[string]int64{
		"id":    counter.ID,
		"count": counter.Count,
	})
	responseBody := bytes.NewBuffer(postBody)

	requestURL := "http://localhost:8080/counters"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}
