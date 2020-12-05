package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/teramont/go-lab-3/server/storage"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type ApiServer struct {
	Port           HttpPortNumber
	StorageHandler storage.HttpHandlerFunc
	server         *http.Server
}

func NewServer(Port HttpPortNumber, StorageHandler storage.HttpHandlerFunc) ApiServer {
	return ApiServer{Port: Port, StorageHandler: StorageHandler, server: nil}
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *ApiServer) Start() error {
	if s.StorageHandler == nil {
		return fmt.Errorf("HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/machines", s.StorageHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *ApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
