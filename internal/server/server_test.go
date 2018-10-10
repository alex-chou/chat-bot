package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMethodHandler(t *testing.T) {
	assert := assert.New(t)
	server := NewServer()
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
