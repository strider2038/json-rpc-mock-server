package app

import (
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/metrics"
	"github.com/creachadair/jrpc2/server"
	"log"
	"net"
	"os"
)

type unixServer struct {
	config     Config
	serviceMap jrpc2.Assigner
}

func NewUnixSocketServer(config Config) Server {
	return &unixServer{
		config:     config,
		serviceMap: serverMethods(),
	}
}

func (s *unixServer) Run() error {
	listener, err := net.Listen("unix", s.config.UnixSocket)
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr, "[unix server] ", log.LstdFlags|log.Lshortfile)
	logger.Printf("starting unix server on socket %s...", s.config.UnixSocket)

	return server.Loop(listener, s.serviceMap, &server.LoopOptions{
		Framing: channel.Line,
		ServerOptions: &jrpc2.ServerOptions{
			Logger:    logger,
			Metrics:   metrics.New(),
			AllowPush: false,
		},
	})
}
