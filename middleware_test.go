package tinyserver

import (
	"net/http"
	"testing"
)

func TestNewMiddleware(t *testing.T) {
	middleware, err := NewMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	if middleware == nil {
		t.Error("Expected NewMiddleware() to return a valid instance of Middleware, got", middleware)
	}
	if err != nil {
		t.Error("Expected NewMiddleware() call to return a nil error, got", err)
	}
}
