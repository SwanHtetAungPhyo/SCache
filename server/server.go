package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/dto"
	"github.com/SwanHtetAungPhyo/Scache/model"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)


const (
	GET = "GET"
	POST  = "POST"
	EVICT = "EVICT"
)


type TCPServer struct {
	Listener net.Listener
	ScacheArray   *model.LRUCache
	RequestCount int
	mu sync.RWMutex
}

func NewScacheServer(config *Config) (*TCPServer, error) {
	scacheArray := model.NewLRUCache(config.Capcity)

	listener, err := net.Listen("tcp",config.Port)
	if err != nil {
		log.Fatalf("Error in listening to the port: %v", err)
		return nil, err
	}
	sServer := &TCPServer{
		Listener:  listener,
		ScacheArray: scacheArray,
		RequestCount: 0,
	}
	go sServer.Start(config.Port)

	return sServer, nil 
}

func (server *TCPServer) Start(port string) {
	for {
		connection, err := server.Listener.Accept()
		connection.Write([]byte(fmt.Sprintf("Cache on Port :%v", port)))
		if err != nil {
			utils.LogMessage(constants.ERROR, err.Error())
			continue
		}
		go server.handleClient(connection)
	}
}

// handleClient manages client connections and requests
func (server *TCPServer) handleClient(connection net.Conn) {
	defer connection.Close()
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		var request dto.Request
		if err := json.Unmarshal(scanner.Bytes(), &request); err != nil {
			connection.Write([]byte("Invalid JSON format\n"))
			continue
		}
		server.mu.Lock()
		server.RequestCount++
        fmt.Printf("Server has been called %d times\n", server.RequestCount)
		server.mu.Unlock()


		switch request.Command {
		case "SET":
			server.handleSet(connection, request)
		case "GET":
			server.handleGet(connection, request)
		default:
			connection.Write([]byte("400 Bad Request\n"))
		}
	}
}


func (server *TCPServer) handleSet(connection net.Conn, request dto.Request) {
	server.ScacheArray.Set(request.Key, request.Value, time.Duration(request.Expiration))

	response := map[string]interface{}{
		"status":  200,
		"message": "Key set successfully",
	}

	responseData, _ := json.Marshal(response)
	connection.Write(append(responseData, '\n')) 
}

func (server *TCPServer) handleGet(connection net.Conn, request dto.Request) {
	value, exists := server.ScacheArray.Get(request.Key)
	if !exists {

		response := map[string]interface{}{
			"status":  404,
			"message": "Key not found",
		}
		responseData, _ := json.Marshal(response)
		connection.Write(append(responseData, '\n'))  
		return
	}


	response := map[string]interface{}{
		"status":  200,
		"value":   value,
	}
	responseData, _ := json.Marshal(response)
	connection.Write(append(responseData, '\n'))  
}
