package main

import (
	"github.com/SwanHtetAungPhyo/Scache/server"
	"github.com/swanhtetaungphyo/Scache/model"
)

func main() {
	var cache model.Scache
	tcpServer, err := server.NewTCPServer(":8000", &cache)
	if err != nil {
		return
	}
	tcpServer.Start()
	
}