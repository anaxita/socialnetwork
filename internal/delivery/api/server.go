// Package api contains objects which are directly used within the REST server
package api

import (
	"context"
	"log"
	"net/http"
	"synergycommunity/internal/delivery/api/middleware"

	"github.com/gorilla/mux"
)

// Server is a structure which contains everything needed for the Community REST server.
type Server struct {
	srv *http.Server
	r   *mux.Router
	gh  http.Handler
	m   *middleware.M
}

// NewServer instantiates a new Server object.
func NewServer(port string, gh http.Handler, m *middleware.M) *Server {
	r := mux.NewRouter()

	s := Server{
		srv: &http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
		r:  r,
		gh: gh,
		m:  m,
	}

	return &s
}

func (s *Server) setGraphQLRoutes() {
	s.r.Handle("/v1/graphql", s.gh)
}

// Start prepares all routes required and listens to the incoming requests.
func (s *Server) Start() error {
	s.r.Use(s.m.Cors.Handler)    // use Cors Middleware for all requests.
	s.r.Use(s.m.Session.Handler) // use Session Middleware for all requests.
	s.r.Use(s.m.Auth.Handler)    // use Auth Middleware for all requests.
	s.setGraphQLRoutes()

	log.Println("Server started at", s.srv.Addr)

	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
