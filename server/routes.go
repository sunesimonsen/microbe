package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (s *Server) setupRoutes() {
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.Logger)

	s.router.Get("/", s.indexHandler)
	s.router.Get("/anchors", s.anchorsHandler)
	s.router.Get("/buttons", s.buttonsHandler)
	s.router.Get("/cards", s.cardsHandler)
	s.router.Get("/colors", s.colorsHandler)
	s.router.Get("/typography", s.typographyHandler)
	s.router.Get("/spacing", s.spacingHandler)
	s.router.Get("/theming", s.themingHandler)
	s.router.Get("/inputs", s.inputsHandler)
	s.router.Get("/checkboxes", s.checkboxesHandler)
	s.router.Get("/radios", s.radiosHandler)
	s.router.Get("/selects", s.selectsHandler)

	s.router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
