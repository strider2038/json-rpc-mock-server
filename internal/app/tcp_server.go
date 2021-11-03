package app

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/metrics"
	"github.com/creachadair/jrpc2/server"
)

type tcpServer struct {
	config     Config
	serviceMap jrpc2.Assigner
}

func NewTCPServer(config Config) Server {
	return &tcpServer{
		config:     config,
		serviceMap: serverMethods(),
	}
}

func (s *tcpServer) Run() error {
	host := fmt.Sprintf(":%d", s.config.Port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	acc := server.NetAccepter(listener, channel.Line)

	logger := log.New(os.Stderr, "[tcp server] ", log.LstdFlags|log.Lshortfile)
	logger.Printf("starting TCP server on port %d...", s.config.Port)

	return server.Loop(acc, server.Static(s.serviceMap), &server.LoopOptions{
		ServerOptions: &jrpc2.ServerOptions{
			Logger:    jrpc2.StdLogger(logger),
			Metrics:   metrics.New(),
			AllowPush: false,
		},
	})
}
