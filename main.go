package main

import (
	"github.com/freedommmoto/rover_control/api"
	"log"
)

func main() {
	server := api.NewServer()
	err := server.Start(getServerAddress())
	if err != nil {
		log.Fatal("can't start server with gin "+getServerAddress(), err)
	}
}

func getServerAddress() string {
	return "0.0.0.0:8884"
}
