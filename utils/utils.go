package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SwanHtetAungPhyo/Scache/constants"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
)

func LogFileConfig() *os.File {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		LogMessage(constants.FATAL, err.Error())
		return nil
	}

	filePath := filepath.Join(homeDir, "Desktop", "server.log")
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

func CurrentFunction() string {
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

func SeverErrorTracker(functionName string, err error) string {
	return fmt.Sprintf("%v : %v", functionName, err.Error())
}

func RequestToServer(request map[string]interface{}, port string) interface{} {
	connection, err := net.Dial("tcp", PortString(port))
	if err != nil {
		return err
	}

	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)

	jsonrquest, _ := json.Marshal(request)
	_, _ = connection.Write(jsonrquest)
	_, _ = connection.Write([]byte("\n"))

	message, _ := bufio.NewReader(connection).ReadString('\n')
	return message
}

func PortString(port string) string {
	return fmt.Sprintf("localhost%v", port)
}
