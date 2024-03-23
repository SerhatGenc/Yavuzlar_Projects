package errors

import (
	"fmt"
	"log"
	"os"
	"time"
)

const logFileName = "Logs/error.log"

func Logger(err error, id int) {
	fmt.Println(err.Error())
	logFile, err2 := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.Ldate|log.Ltime)

	now := time.Now()

	logger.Printf("[%s] ID: %d - Error: %v\n", now.Format("2006-01-02 15:04:05"), id, err)
}
