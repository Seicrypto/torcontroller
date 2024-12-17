package logger

import (
	"log"
	"os"
	"sync"
)

var (
	instance    *log.Logger
	once        sync.Once
	logFilePath = "/var/log/torcontroller.log"
)

func GetLogger() *log.Logger {
	once.Do(func() {
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to create log file: %v", err)
		}

		instance = log.New(logFile, "TORCONTROLLER: ", log.Ldate|log.Ltime|log.Lshortfile)
	})
	return instance
}
