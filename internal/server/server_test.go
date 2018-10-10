package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alex-chou/chat-bot/internal/backend/backendmocks"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	backend *backendmocks.Backend
}

func initialize() (*Server, *mocks) {
	backend := new(backendmocks.Backend)
	return NewServer(backend), &mocks{
		backend: backend,
	}
}

func TestNewMethodHandler(t *testing.T) {
	assert := assert.New(t)
	server, _ := initialize()
	handler := NewWrappedHandler(http.MethodGet, server.Health)

	t.Run("correct method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		assert.Nil(err)

		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusOK, rr.Code)
	})

	t.Run("incorrect method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodHead, "/health", nil)
		assert.Nil(err)

		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusNotFound, rr.Code)
		assert.Equal("404 page not found\n", rr.Body.String())
	})
}
