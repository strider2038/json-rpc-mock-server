package app

import (
	"fmt"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/creachadair/jrpc2/metrics"
	"github.com/creachadair/jrpc2/server"
	"log"
	"net/http"
	"os"
)

type httpServer struct {
	config     Config
	serviceMap jrpc2.Assigner
}

func NewHTTPServer(config Config) Server {
	return &httpServer{
		config:     config,
		serviceMap: serverMethods(),
	}
}

func (s *httpServer) Run() error {
	logger := log.New(os.Stderr, "[http server] ", log.LstdFlags|log.Lshortfile)
	local := server.NewLocal(s.serviceMap, &server.LocalOptions{
		Server: &jrpc2.ServerOptions{
			Logger:  logger,
			Metrics: metrics.New(),
		},
	})
	http.Handle("/rpc", jhttp.NewBridge(local.Client))

	logger.Printf("starting HTTP server on port %d...", s.config.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), nil)
}
