package server

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) setupRoutes() {
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.Logger)

	s.router.Get("/", IndexHandler)
	s.router.Get("/{category}/{page}", DocsHandler)
	s.router.Get("/search", SearchHandler)

	s.router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
