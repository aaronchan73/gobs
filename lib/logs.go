package gobs

import (
	"fmt"
	"time"
)

// Log is a single message with a timestamp
type Log struct {
	ID        int64     `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// CreateLog creates a Log
func CreateLog(id int64, message string) Log {
	log := Log{
		id,
		time.Now(),
		message,
	}

	return log
}

// PrintLog prints a log
func PrintLog(log Log) {
	logString := "[" + log.Timestamp.Format(time.TimeOnly) + "] " + log.Message
	fmt.Println(logString)
}
