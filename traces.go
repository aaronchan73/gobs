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

// createSpan creates a Span
func createSpan(event string, duration time.Duration) Span {
	span := Span{
		duration,
		event,
	}

	return span
}

// printSpan prints a Span
func printSpan(span Span) {
	spanString := fmt.Sprintf("%s: %d ms\n", span.event, span.duration.Milliseconds())
	fmt.Println(spanString)
}

// createTrace creates a Trace
func createTrace() *Trace {
	trace := Trace{
		make([]Span, 0),
		&sync.RWMutex{},
	}

	return &trace
}

// updateTrace updates an existing Trace
func (trace *Trace) updateTrace(span Span) {
	trace.lock.Lock()
	defer trace.lock.Unlock()

	trace.spans = append(trace.spans, span)
}

// printTrace prints an existing Trace
func (trace *Trace) printTrace() {
	trace.lock.RLock()
	defer trace.lock.RUnlock()

	for _, span := range trace.spans {
		printSpan(span)
	}
}
