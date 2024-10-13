package main

import (
	"github.com/SwanHtetAungPhyo/Scache/server"
	"log"
	"net"
	"time"
)

func main() {
	cacheConfig, err := server.NewCofig(
		server.WithPort(":9000"),
		server.WithCapacity(200),
		server.WithExpiration(10*time.Minute),
	)
	// Adjust the import path accordingly
	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}
	s, err := server.NewScacheServer(cacheConfig)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer func(Listener net.Listener) {
		err := Listener.Close()
		if err != nil {

		}
	}(s.Listener)

	log.Println("Server is running on port:", cacheConfig.Port)
	select {} // Keep the server running
}
