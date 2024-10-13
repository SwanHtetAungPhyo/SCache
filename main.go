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
	
	request := map[string]interface{}{
        "command":    "SET",
        "key":        "username",
        "value":      "john_doe",
        "expiration": 60000000000, // 1 minute in nanoseconds
    }

	RequestToServer(request, cacheConfig.Port)
	select{}
}

func RequestToServer(request map[string]interface{},port string) interface{}{
	connection,err := net.Dial("tcp",PortString(port) )
	if err != nil{
		return err
	}
	
	defer connection.Close()

	jsonrquest, _ := json.Marshal(request)
	connection.Write(jsonrquest)
	connection.Write([]byte("\n"))

	message, _ := bufio.NewReader(connection).ReadString('\n')
	return message
}


func PortString(port string) string{
	return fmt.Sprintf("localhost%v", port)
}