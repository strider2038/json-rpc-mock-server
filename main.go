package main

import (
	"flag"
	"json-rpc-mock-server/internal/app"
	"log"
)

func main() {
	protocol := flag.String("protocol", "tcp", "server protocol")
	flag.Parse()

	config := app.LoadConfigFromEnvironment()
	var server app.Server
	switch *protocol {
	case "tcp":
		server = app.NewTCPServer(config)
	case "http":
		server = app.NewHTTPServer(config)
	default:
		log.Fatalf("unsupported protocol: %s", *protocol)
	}

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
