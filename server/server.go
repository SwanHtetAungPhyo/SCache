package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/dto"
	"github.com/SwanHtetAungPhyo/Scache/model"
	"github.com/SwanHtetAungPhyo/Scache/utils"
	"log"
	"net"
	"sync"
	"time"
)

const (
	GET   = "GET"
	POST  = "POST"
	EVICT = "EVICT"
)

type TCPServer struct {
	Listener     net.Listener
	ScacheArray  *model.LRUCache
	RequestCount int
	mu           sync.RWMutex
}

func NewScacheServer(config *Config) (*TCPServer, error) {
	scacheArray := model.NewLRUCache(config.Capcity)

	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("Error in listening to the port: %v", err)
		return nil, err
	}
	sServer := &TCPServer{
		Listener:     listener,
		ScacheArray:  scacheArray,
		RequestCount: 0,
	}
	go sServer.Start(config.Port)

	return sServer, nil
}

func (server *TCPServer) Start(port string) {
	for {
		connection, err := server.Listener.Accept()
		_, _ = connection.Write([]byte(fmt.Sprintf("Cache on Port :%v", port)))
		if err != nil {
			utils.LogMessage(constants.ERROR, err.Error())
			continue
		}
		go server.handleClient(connection)
	}
}

// handleClient manages client connections and requests
func (server *TCPServer) handleClient(connection net.Conn) {
	defer func() {
		_ = connection.Close()
	}()

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		var requests dto.Requests // Change here to handle multiple requests
		if err := json.Unmarshal(scanner.Bytes(), &requests); err != nil {
			_, _ = connection.Write([]byte("Invalid JSON format\n"))
			continue
		}

		server.mu.Lock()
		server.RequestCount += len(requests.Requests) // Count total requests received
		fmt.Printf("Server has been called %d times\n", server.RequestCount)
		server.mu.Unlock()

		for _, request := range requests.Requests { // Loop through each request
			switch request.Command {
			case "SET":
				server.handleSet(connection, request)
			case "GET":
				server.handleGet(connection, request)
			default:
				_, _ = connection.Write([]byte("400 Bad Request\n"))
			}
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
	_, _ = connection.Write(append(responseData, '\n'))
}

func (server *TCPServer) handleGet(connection net.Conn, request dto.Request) {
	value, exists := server.ScacheArray.Get(request.Key)
	if !exists {

		response := map[string]interface{}{
			"status":  404,
			"message": "Key not found",
		}
		responseData, _ := json.Marshal(response)
		_, _ = connection.Write(append(responseData, '\n'))
		return
	}

	response := map[string]interface{}{
		"status": 200,
		"value":  value,
	}
	responseData, _ := json.Marshal(response)
	_, err := connection.Write(append(responseData, '\n'))
	if err != nil {
		return
	}
}
