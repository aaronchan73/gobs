package gobs

import (
	"fmt"
	"time"
)

type Log struct {
	timestamp time.Time
	message   string
}

// createLog creates a Log
func createLog(message string) Log {
	log := Log{
		time.Now(),
		message,
	}

	return log
}

// printLog prints a log
func printLog(log Log) {
	logString := log.timestamp.Format(time.TimeOnly) + log.message
	fmt.Println(logString)
}
