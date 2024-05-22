package gobs

import (
	"fmt"
	"time"
)

// Log is a single message with a timestamp
type Log struct {
	timestamp time.Time
	message   string
}

// CreateLog creates a Log
func CreateLog(message string) Log {
	log := Log{
		time.Now(),
		message,
	}

	return log
}

// PrintLog prints a log
func PrintLog(log Log) {
	logString := log.timestamp.Format(time.TimeOnly) + log.message
	fmt.Println(logString)
}
