package main

import (
	"github.com/freedommmoto/rover_control/api"
	"log"
	"os"
)

func main() {
	server := api.NewServer()
	err := server.Start(getServerAddress(), getFilePart())
	if err != nil {
		log.Fatal("can't start server with gin "+getServerAddress(), err)
	}
}

func getServerAddress() string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return "0.0.0.0:7472"
}

func getFilePart() string {
	if os.Getenv("FILE_PART") != "" {
		return os.Getenv("FILE_PART")
	}
	return "../instructions_file/instructions.txt"
}
