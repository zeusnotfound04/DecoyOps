package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time" // Added for timestamp
)

type LogEntry struct {
	TechniqueID   string `json:"technique_id"`
	TechniqueName string `json:"technique_name"`
	Timestamp     string `json:"timestamp"` // Added for timestamp
	Message       string `json:"message"`
}

func InitLogger() {
	os.MkdirAll("output", os.ModePerm)
}

func Log(techniqueID, techniqueName, message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05") // Get current time
	entry := LogEntry{
		TechniqueID:   techniqueID,
		TechniqueName: techniqueName,
		Timestamp:     currentTime, // Add timestamp to log entry
		Message:       message,
	}

	// Added some basic styling (colors) for CLI output
	// You might need a library for more advanced styling
	fmt.Printf("\\033[34m[%s]\\033[0m \\033[32m%s\\033[0m - \\033[33m%s\\033[0m - %s\\n", currentTime, techniqueID, techniqueName, message)

	f, _ := os.OpenFile("output/output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	enc := json.NewEncoder(f)

	enc.Encode(entry)
}
