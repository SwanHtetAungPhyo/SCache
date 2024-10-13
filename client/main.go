package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/server"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)

func main() {
	// Configure cache server
	cacheConfig, err := server.NewCofig(
		server.WithPort(":9000"),
		server.WithCapacity(200),
		server.WithExpiration(10 * time.Minute),
	)
	if err != nil {
		panic("Cache Server Configuration Failed: " + err.Error())
	}

	_, err = server.NewScacheServer(cacheConfig)
	if err != nil {
		utils.LogMessage(constants.ERROR, err.Error())
		return
	}

	// Connect to the cache server
	connection, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer connection.Close()

	// Prepare request object
	requestObj := map[string]interface{}{
		"command":    "SET",
		"key":       "user123",
		"value":     "John Doe",
		"expiration": time.Now().Add(10 * time.Minute).UnixNano(), // Use current time for expiration
	}

	// Marshal request to JSON
	jsonRequest, err := json.Marshal(requestObj)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return
	}

	// Send request to server
	_, err = connection.Write(append(jsonRequest, '\n'))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// Read response from server
	message, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Unmarshal the response into a Go object
	var responseObj map[string]interface{}
	err = json.Unmarshal([]byte(message), &responseObj)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	// Save the response to a JSON file
	responseFile, err := os.Create("server_response.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer responseFile.Close()

	// Encode the response object into the file
	encoder := json.NewEncoder(responseFile)
	encoder.SetIndent("", "  ") // Optional: pretty-print the JSON
	err = encoder.Encode(responseObj)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}

	// Print the server response to console
	fmt.Print("Server response saved to server_response.json")
}
