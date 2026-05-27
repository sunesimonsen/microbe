package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (s *Server) setupRoutes() {
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.Logger)

	s.router.Get("/", s.indexHandler)
	s.router.Get("/anchor", s.anchorHandler)
	s.router.Get("/button", s.buttonHandler)
	s.router.Get("/card", s.cardHandler)
	s.router.Get("/colors", s.colorsHandler)
	s.router.Get("/typography", s.typographyHandler)
	s.router.Get("/spacing", s.spacingHandler)
	s.router.Get("/theming", s.themingHandler)
	s.router.Get("/input", s.inputHandler)
	s.router.Get("/checkbox", s.checkboxHandler)
	s.router.Get("/switch", s.switchHandler)
	s.router.Get("/radio", s.radioHandler)
	s.router.Get("/select", s.selectHandler)
	s.router.Get("/range", s.rangeHandler)

	s.router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
