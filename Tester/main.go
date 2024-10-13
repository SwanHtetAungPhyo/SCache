package main

import (
	"github.com/SwanHtetAungPhyo/Scache/server"
	
)

func main() {
	var cache .Scache
	tcpServer, err := server.NewTCPServer(":8000", &cache)
	if err != nil {
		return
	}
	tcpServer.Start()
	
}