package gobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

// UpdateLog updates an existing log
func UpdateLog(log Log) {
	jsonBody, _ := json.Marshal(map[string]interface{}{
		"id":        log.ID,
		"timestamp": log.Timestamp,
		"message":   log.Message,
	})
	responseBody := bytes.NewBuffer(jsonBody)

	requestURL := os.Getenv("COLLECTOR_ADDRESS") + "/logs"
	if _, err := http.Post(requestURL, "application/json", responseBody); err != nil {
		panic(err)
	}
}

// PrintLog prints a log
func PrintLog(log Log) {
	logString := "[" + log.Timestamp.Format(time.TimeOnly) + "] " + log.Message
	fmt.Println(logString)
}
