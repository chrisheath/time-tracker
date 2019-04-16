package logger

import (
	"go/build"
	"log"
	"os"
)

// Write line to log file
func Write(logString []string) {
	logPath := build.Default.GOPATH + "/src/github.com/chrisheath/time-tracker/log.txt"

	var file, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	var logToFile string

	for i := 0; i < len(logString); i++ {
		logToFile += string(logString[i]) + " "
	}

	log.SetOutput(file)
	log.Printf(logToFile)
}
