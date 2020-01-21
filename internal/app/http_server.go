package app

import (
	"fmt"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/creachadair/jrpc2/metrics"
	"github.com/creachadair/jrpc2/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"regexp"
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

	router := mux.NewRouter()
	router.Handle("/rpc", jhttp.NewBridge(local.Client))

	if s.config.BearerToken != "" {
		router.Use(s.bearerAuthorization)
		logger.Printf("server requires authorization header with token '%s'", s.config.BearerToken)
	}

	logger.Printf("starting HTTP server on port %d...", s.config.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), router)
}

func (s *httpServer) bearerAuthorization(next http.Handler) http.Handler {
	re := regexp.MustCompile("^Bearer (.+)$")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, exists := r.Header["Authorization"]
		if !exists || len(header) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		matches := re.FindStringSubmatch(header[0])
		if len(matches) < 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token := matches[1]
		if token != s.config.BearerToken {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
