package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router chi.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() (*Server, error) {
	router := chi.NewRouter()

	s := &Server{
		router: router,
	}

	s.setupRoutes()

	return s, nil
}
