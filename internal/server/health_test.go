package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert := assert.New(t)
	server := NewServer()
	handler := NewWrappedHandler(http.MethodGet, server.Health)

	t.Run("successful call", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/health", nil)
		assert.Nil(err)

		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusOK, rr.Code)
		assert.Equal("Hello! The chatbot is healthy.\n", rr.Body.String())
	})
}
