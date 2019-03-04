package app

import (
	"os"
	"strconv"
)

type Config struct {
	Port         uint16
	ThreadsCount uint16
}

func LoadConfigFromEnvironment() Config {
	return Config{
		Port:         getPort(),
		ThreadsCount: getThreadsCount(),
	}
}

func getPort() uint16 {
	unparsedPort := os.Getenv("JSON_RPC_PORT")
	port, _ := strconv.ParseUint(unparsedPort, 10, 32)

	if port == 0 {
		port = 4000
	}

	return uint16(port)
}

func getThreadsCount() uint16 {
	unparsedThreads := os.Getenv("JSON_RPC_THREADS")
	threads, _ := strconv.ParseUint(unparsedThreads, 10, 32)

	return uint16(threads)
}
