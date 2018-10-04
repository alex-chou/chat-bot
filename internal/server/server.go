package server

import (
	"net/http"
	"time"
)

const (
	// timeoutDuration is the default timeout duration for every handler.
	timeoutDuration = 5 * time.Second
	// timeoutMessage is the default timeout message for every handler.
	timeoutMessage = "{\"error\": \"endpoint timeout\"}"
)

// Server is a chatbot server.
type Server struct {
	*http.ServeMux
}

// NewServer creates a new server.
func NewServer() *Server {
	server := &Server{
		ServeMux: http.NewServeMux(),
	}
	server.HandleFunc("/health", server.Health)
	return server
}

// HandleFunc wraps the handler and passes it to the embedded ServeMux.
func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	wrappedHandler := http.TimeoutHandler(http.HandlerFunc(handler), timeoutDuration, timeoutMessage)
	s.ServeMux.Handle(pattern, wrappedHandler)
}
