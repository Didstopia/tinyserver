package tinyserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	server, err := NewServer()
	if server == nil {
		t.Error("Expected NewServer() to return a valid instance of Server, got", server)
	}
	if err != nil {
		t.Error("Expected NewServer() call to return a nil error, got", err)
	}
}

func TestAddRoute(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Error("Expected NewServer() call to return a nil error, got", err)
	}
	handler := func(w http.ResponseWriter, r *http.Request) {}
	err = server.AddRoute("/", handler)
	if err != nil {
		t.Error("Expected AddRoute() call to return a nil error, got", err)
	}
}

func TestAddMiddleware(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Error("Expected NewServer() call to return a nil error, got", err)
	}
	middleware := JSONMiddleware
	err = server.AddMiddleware(middleware)
	if err != nil {
		t.Error("Expected AddMiddleware() call to return a nil error, got", err)
	}
}

func TestListenAndClose(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Error("Expected NewServer() call to return a nil error, got", err)
	}
	// FIXME: We need to find a way to test the synchronous "Listen()" method
	server.ListenAsync("8080")
	resp := httptest.NewRecorder()
	req, reqErr := http.NewRequest("GET", "/asdf", nil)
	if reqErr != nil {
		t.Error("Request failed:", reqErr)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Error("Failed to read response body:", err, body)
	} else {
		if !strings.Contains(string(body), "404") {
			t.Error("Expected GET /asdf to return a 404, got", string(body))
		}
	}
	err = server.Close()
	if err != nil {
		t.Error("Expected Close() call to return a nil error, got", err)
	}
}
