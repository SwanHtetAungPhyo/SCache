package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/dto"
	"github.com/SwanHtetAungPhyo/Scache/model"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)



type TCPServer struct {
	Listener net.Listener
	ScacheArray   *model.LRUCache
	RequestCount int
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
		server.RequestCount++
        fmt.Printf("Server has been called %d times\n", server.RequestCount)
		newLRU := model.NewLRUCache(10)
		switch request.Command {
		case "SET":
			newLRU.Set(request.Key,request.Value,time.Duration(request.Expiration))
			connection.Write([]byte("200\n"))
		case "GET":
			value, codition := newLRU.Get(request.Key)
			if !codition{
				connection.Write([]byte("500"))
			}
			response, _ := json.Marshal(value)
			connection.Write(append(response, '\n'))
		default:
			connection.Write([]byte("400 Bad Request\n"))
		}
	}
}
