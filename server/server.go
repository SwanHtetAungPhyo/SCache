package server

import (
	"bufio"
	"encoding/json"
	"log"
	"net"

	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/dto"
	"github.com/SwanHtetAungPhyo/Scache/model"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)


type TCPServer struct {
	Listener net.Listener
	Scache   *model.Scache
	RequestCount int
}

// NewTCPServer initializes a new TCP server
func NewTCPServer(port string, scache *model.Scache) (*TCPServer, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error in listening to the port: %v", err)
		return nil, err
	}
	return &TCPServer{Listener: listener, Scache: scache}, nil
}

// Start begins accepting client connections
func (server *TCPServer) Start() {
	for {
		connection, err := server.Listener.Accept()
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
        // fmt.Printf("Server has been called %d times\n", server.RequestCount)
		// switch request.Command {
		// case "SET":
		// 	server.Scache.Set(request.Key, request.Value)
		// 	connection.Write([]byte("200\n"))
		// case "GET":
		// 	value, _ := server.Scache.Get(request.Key)
		// 	response, _ := json.Marshal(value)
		// 	connection.Write(append(response, '\n'))
		// default:
		// 	connection.Write([]byte("400 Bad Request\n"))
		// }
	}
}
