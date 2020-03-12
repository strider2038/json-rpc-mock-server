package main

import (
	"json-rpc-mock-server/internal/app"
	"log"
	"strings"
)

func main() {
	config := app.LoadConfigFromEnvironment()
	var server app.Server
	protocol := strings.ToLower(config.Protocol)
	switch protocol {
	case "tcp":
		server = app.NewTCPServer(config)
	case "http":
		server = app.NewHTTPServer(config)
	case "unix":
		server = app.NewUnixSocketServer(config)
	default:
		log.Fatalf("unsupported protocol: %s", protocol)
	}

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
