package handler

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

		"github.com/SwanHtetAungPhyo/Scache/utils"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

func HandleConnections(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error in accepting the connection: %v", err)
			continue
		}
		connnectionString := conn.RemoteAddr().String()
		log.Printf("Connection came from port: %v", connnectionString)
		go handlerClient(conn) 
	}
}


func handlerClient(connection net.Conn){
	byteValue , err := os.ReadFile("../guide.txt")
	connection.Write(byteValue)
	if err != nil{
		utils.LogMessage(ERROR, err.Error())
		return
	}

	
	defer func ()  {
		if err := connection.Close(); err != nil{
			utils.LogMessage(ERROR, err.Error())
		}
	}()

}

func CommandParser(connection net.Conn) {
	scanner := bufio.NewScanner(connection)

	for scanner.Scan(){
		request := scanner.Text()
		fmt.Printf("Received request : %v from Client Address: %v\n", request, connection.RemoteAddr().String())
		
	}
}

func connnectionWriter(connection net.Conn, message string){
	_, err := connection.Write([]byte(message))
	if err != nil{
		utils.LogMessage(ERROR, err.Error())
		return
	}	
}