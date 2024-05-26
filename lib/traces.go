package gobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Span is a single event with a duration
type Span struct {
	ID       int64         `json:"id"`
	Duration time.Duration `json:"duration"`
	Event    string        `json:"event"`
}

// Trace is a collection of Spans
type Trace struct {
	ID    int64  `json:"id"`
	Spans []Span `json:"spans"`
}

// CreateSpan creates a Span
func CreateSpan(id int64, event string, duration time.Duration) Span {
	span := Span{
		id,
		duration,
		event,
	}

	return span
}

// PrintSpan prints a Span
func PrintSpan(span Span) {
	spanString := fmt.Sprintf("%s: %d ms", span.Event, span.Duration.Milliseconds())
	fmt.Println(spanString)
}

// CreateTrace creates a Trace
func CreateTrace(id int64) *Trace {
	trace := Trace{
		id,
		make([]Span, 0),
	}

	return &trace
}

// UpdateTrace updates an existing Trace
func (trace *Trace) UpdateTrace(span Span) {
	trace.Spans = append(trace.Spans, span)

	jsonBody, _ := json.Marshal(map[string]interface{}{
		"id":      trace.ID,
		"buckets": trace.Spans,
	})
	responseBody := bytes.NewBuffer(jsonBody)

	requestURL := os.Getenv("COLLECTOR_ADDRESS") + "/traces"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}

// PrintTrace prints an existing Trace
func (trace *Trace) PrintTrace() {
	for _, span := range trace.Spans {
		PrintSpan(span)
	}
}
