package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/SwanHtetAungPhyo/Scache/constants"
)



func LogFileConfig() *os.File {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		LogMessage(constants.FATAL,err.Error())
		return nil 
	}

	filePath := filepath.Join(homeDir, "Desktop","server.log")
	logfile, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error in opening the log file: %v", err)
		return nil
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logfile
}

func LogMessage(level string, message string) {
	log.Printf("[%s] %s", level, message)
}

func CurrentFunction() string{
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}


func InitLog() *os.File {
	logfile := LogFileConfig()
	if logfile == nil {
		fmt.Println("Returned log file is nil")
		return nil
	}
	return logfile
}


func CloseLogFile(logfile *os.File) {
	if err := logfile.Close(); err != nil {
		fmt.Println("Cannot close the log file:", err)
	}
}

func SeverErrorTracker(functionName string,err error) string {
	return fmt.Sprintf("%v : %v", functionName, err.Error())
}