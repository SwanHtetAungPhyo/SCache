package main

import (
	"github.com/SwanHtetAungPhyo/Scache/server"
	"time"
	"github.com/SwanHtetAungPhyo/Scache/utils"
	"github.com/SwanHtetAungPhyo/Scache/constants"
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

	
	select{}
}