package tinyserver

import "net/http"

// MiddlewareFunc is an abstraction for global or path level middleware function
type MiddlewareFunc func(next http.Handler) http.Handler

// Middleware is an interface for MiddlewareFunc
type Middleware interface {
	Middleware(handler http.Handler) http.Handler
}

// Middleware implements the Middleware interface for MiddlewareFunc
func (middleware MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return middleware(handler)
}

// NewMiddleware creates a new Middleware and returns it
func NewMiddleware(handler func(w http.ResponseWriter, r *http.Request)) (Middleware, error) {
	// TODO: Isn't this kind of dumb and useless for our use case?
	middlewareHandler := http.HandlerFunc(handler)
	var middleware MiddlewareFunc = func(next http.Handler) http.Handler {
		return middlewareHandler
	}
	return middleware, nil
}

// JSONMiddleware returns a request handler for application/json content
func JSONMiddleware(next http.Handler) http.Handler {
	// TODO: If we're sticking with "NewMiddleware", shouldn't we return "Middleware" here instead of "http.Handler"?
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
