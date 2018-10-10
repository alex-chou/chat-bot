package server

import (
	"fmt"
	"net/http"
)

// Health responds with a default message.
func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello! The chatbot is healthy.")
}
