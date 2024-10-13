package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/server"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)
func main(){
	cacheConfig, err := server.NewCofig(
		server.WithPort(":9000"),
		server.WithCapacity(200),
		server.WithExpiration(10 * time.Minute),
	)
	if err != nil {
		panic("Cache Server Configuration Failed")
	}
	_, err = server.NewScacheServer(cacheConfig)
	if err != nil{
		 utils.LogMessage(constants.ERROR, err.Error())
	}	

		
	connection,err := net.Dial("tcp","localhost:9000" )
	if err != nil{
		return 
	}
	
	defer connection.Close()
	requestObj := map[string]interface{}{
		"command":    "SET",
		"key":       "user123",
		"value":     "John Doe",
		"expiration": 1697270400,
	}

	jsonrquest, _ := json.Marshal(requestObj)
	connection.Write(jsonrquest)
	connection.Write([]byte("\n"))

	message, _ := bufio.NewReader(connection).ReadString('\n')
	fmt.Print(message)

	select{}
}