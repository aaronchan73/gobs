package gobs

import (
	"fmt"
	"sync"
	"time"
)

// Span is a single event with a duration
type Span struct {
	duration time.Duration
	event    string
}

// Trace is a collection of Spans
type Trace struct {
	spans []Span
	lock  *sync.RWMutex
}

// CreateSpan creates a Span
func CreateSpan(event string, duration time.Duration) Span {
	span := Span{
		duration,
		event,
	}

	return span
}

// PrintSpan prints a Span
func PrintSpan(span Span) {
	spanString := fmt.Sprintf("%s: %d ms\n", span.event, span.duration.Milliseconds())
	fmt.Println(spanString)
}

// CreateTrace creates a Trace
func CreateTrace() *Trace {
	trace := Trace{
		make([]Span, 0),
		&sync.RWMutex{},
	}

	return &trace
}

// UpdateTrace updates an existing Trace
func (trace *Trace) UpdateTrace(span Span) {
	trace.lock.Lock()
	defer trace.lock.Unlock()

	trace.spans = append(trace.spans, span)
}

// PrintTrace prints an existing Trace
func (trace *Trace) PrintTrace() {
	trace.lock.RLock()
	defer trace.lock.RUnlock()

	for _, span := range trace.spans {
		PrintSpan(span)
	}
}
