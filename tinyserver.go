package tinyserver

import (
	"log"
	"net/http"
)

// Server is a top-level abstraction for the web server
type Server struct {
	// Routes stores each route separately
	Routes []*Route

	// HTTPServer is the primary HTTP server
	HTTPServer *http.Server

	// Router is the primary request router
	Router *http.ServeMux

	// Middlewares stores global middlewares
	Middlewares []Middleware
}

// NewServer creates a new instance of Server and returns a reference to it
func NewServer() (*Server, error) {
	// Create a new instance of the server
	server := &Server{}

	// Create a new request router
	server.Router = http.NewServeMux()

	// Add default middlewares
	err := server.AddMiddleware(JSONMiddleware) // TODO: Instead of adding default middlewares, expose this with something like a "EnableJSON" function?

	// Add default routes
	err = server.AddRoute("/", func(w http.ResponseWriter, r *http.Request) {})

	// Return the server instance and a potential error
	return server, err
}

// AddRoute constructs and stores a new route
func (server *Server) AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) error {
	//route := &Route{Path: path, Handler: http.HandlerFunc(handler)}
	route, err := NewRoute()
	if err != nil {
		log.Panic(err)
	}
	route.Path = path
	route.Handler = http.HandlerFunc(handler)
	server.Routes = append(server.Routes, route)
	return nil
}

// AddMiddleware adds one or more global middleware
func (server *Server) AddMiddleware(middlewares ...MiddlewareFunc) error {
	for _, middleware := range middlewares {
		server.Middlewares = append(server.Middlewares, middleware)
	}
	return nil
}

// Listen starts the HTTP server on the specified port (synchronous/blocking)
func (server *Server) Listen(port string) error {
	server.initialize(port)
	// FIXME: ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is ErrServerClosed
	err := server.HTTPServer.ListenAndServe()
	return err
}

// ListenAsync starts the HTTP server on the specified port (asynchronous/non-blocking)
func (server *Server) ListenAsync(port string) {
	server.initialize(port)
	// FIXME: ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is ErrServerClosed
	go server.HTTPServer.ListenAndServe()
}

// Close the HTTP server and cleanup after ourselves
func (server *Server) Close() error {
	log.Println("Closing server..")
	return server.HTTPServer.Close()
}

// Internal server initialization logic
func (server *Server) initialize(port string) {
	server.HTTPServer = &http.Server{Addr: ":" + port, Handler: server.Router}

	// Setup routes and middlewares
	for _, route := range server.Routes {
		log.Println("Adding route", route.Path)

		// Apply each middleware to this route
		handler := route.Handler
		for _, middleware := range server.Middlewares {
			handler = middleware.Middleware(handler)
		}

		// Add the route with the final handler
		server.Router.Handle(route.Path, handler)
	}

	log.Println("Server listening on port", port)
}
