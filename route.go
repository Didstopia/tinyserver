package tinyserver

import (
	"net/http"
)

// Route matches a request path and calls the handler function
type Route struct {
	// Path for this route
	Path string

	// Handler for this route
	Handler http.Handler
}

// NewRoute creates a new instance of Route and returns a reference to it
func NewRoute() (*Route, error) {
	route := &Route{}
	return route, nil
}
