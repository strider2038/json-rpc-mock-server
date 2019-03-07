package app

import (
	"bitbucket.org/creachadair/jrpc2"
	"bitbucket.org/creachadair/jrpc2/channel"
	"bitbucket.org/creachadair/jrpc2/handler"
	"bitbucket.org/creachadair/jrpc2/metrics"
	jrpcServer "bitbucket.org/creachadair/jrpc2/server"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"time"
)

type JsonRpcServer interface {
	Run() error
}

type SleepDuration struct {
	StartTime int64
	EndTime   int64
}

type tcpServer struct {
	port         uint16
	threadsCount uint16
	serviceMap   jrpc2.Assigner
}

func NewJsonRpcServer(config Config) JsonRpcServer {
	server := &tcpServer{
		port:         config.Port,
		threadsCount: config.ThreadsCount,
	}

	server.serviceMap = handler.Map{
		"sum":      handler.New(server.sum),
		"subtract": handler.New(server.subtract),
		"sleep":    handler.New(server.sleep),
		"reflect":  handler.New(server.reflect),
		"notify":   handler.New(server.notify),
	}

	return server
}

func (server *tcpServer) Run() error {
	host := fmt.Sprintf(":%d", server.port)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}

	logger := logrus.New()
	logger.Infof("Listening for TCP connections at port %d...", server.port)

	writer := logger.Writer()
	defer func() {
		_ = writer.Close()
	}()

	err = jrpcServer.Loop(listener, server.serviceMap, &jrpcServer.LoopOptions{
		Framing: channel.Line,
		ServerOptions: &jrpc2.ServerOptions{
			Logger:      log.New(writer, "[TCP] ", log.Lshortfile),
			Concurrency: int(server.threadsCount),
			Metrics:     metrics.New(),
			AllowPush:   false,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (server *tcpServer) sum(context context.Context, arguments []int) (int, error) {
	var summary int

	for _, value := range arguments {
		summary += value
	}

	return summary, nil
}

func (server *tcpServer) subtract(context context.Context, arguments []int) (int, error) {
	if len(arguments) != 2 {
		return 0, errors.New("there must be exactly two arguments")
	}

	result := arguments[1] - arguments[0]

	return result, nil
}

func (server *tcpServer) sleep(context context.Context, arguments []int) (SleepDuration, error) {
	duration := SleepDuration{}
	duration.StartTime = time.Now().UnixNano()

	sleepingTime := 1000

	if len(arguments) > 0 {
		sleepingTime = arguments[0]
	}

	time.Sleep(time.Duration(sleepingTime) * time.Millisecond)

	duration.EndTime = time.Now().UnixNano()

	return duration, nil
}

func (server *tcpServer) reflect(context context.Context, arguments map[string]interface{}) (map[string]interface{}, error) {
	return arguments, nil
}

func (server *tcpServer) notify(context context.Context, arguments map[string]interface{}) error {
	return nil
}
