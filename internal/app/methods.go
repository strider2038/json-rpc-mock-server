package app

import (
	"context"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"time"
)

func serverMethods() jrpc2.Assigner {
	return handler.Map{
		"ping":     handler.New(ping),
		"sum":      handler.New(sum),
		"subtract": handler.New(subtract),
		"sleep":    handler.New(sleep),
		"reflect":  handler.New(reflect),
		"notify":   handler.New(notify),
	}
}

type SleepDuration struct {
	StartTime int64
	EndTime   int64
}

type SubtractionParameters struct {
	Subtrahend int `json:"subtrahend"`
	Minuend    int `json:"minuend"`
}

func ping(context context.Context) (string, error) {
	return "pong", nil
}

func sum(context context.Context, arguments []int) (int, error) {
	var summary int

	for _, value := range arguments {
		summary += value
	}

	return summary, nil
}

func subtract(context context.Context, arguments SubtractionParameters) (int, error) {
	result := arguments.Minuend - arguments.Subtrahend

	return result, nil
}

func sleep(context context.Context, arguments []int) (SleepDuration, error) {
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

func reflect(context context.Context, arguments map[string]interface{}) (map[string]interface{}, error) {
	return arguments, nil
}

func notify(context context.Context, arguments map[string]interface{}) error {
	return nil
}
