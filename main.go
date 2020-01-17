package main

import (
	"json-rpc-mock-server/internal/app"
	"log"
)

func main() {
	config := app.LoadConfigFromEnvironment()
	server := app.NewJsonRpcServer(config)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
