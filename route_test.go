package tinyserver

import "testing"

func TestNewRoute(t *testing.T) {
	route, err := NewRoute()
	if route == nil {
		t.Error("Expected NewRoute() to return a valid instance of Route, got", route)
	}
	if err != nil {
		t.Error("Expected NewRoute() call to return a nil error, got", err)
	}
}
