package logger

import (
	"log"
	"os"
)

var logFilePath = "/var/log/torcontroller.log"

func CreateLogger() (*log.Logger, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger := log.New(logFile, "TORCONTROLLER: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("========== Starting new session ==========")
	return logger, nil
}
