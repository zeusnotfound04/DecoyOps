package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogEntry struct {
	TechniqueID   string `json:"technique_id"`
	TechniqueName string `json:"technique_name"`
	Message       string `json:"message'`
}

func InitLogger() {
	os.MkdirAll("output" , os.ModePerm)
}


func Log(techniqueID , techniqueName , message string) {
	entry := LogEntry{
		TechniqueID: techniqueID,
		TechniqueName: techniqueName,
		Message: message,
	}


	fmt.Printf("[%s] %s - %s\n" , techniqueID , techniqueName , message)


	f , _ := os.OpenFile("output/output.json" , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	enc := json.NewEncoder(f)

	enc.Encode(entry)
}