package server

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sunesimonsen/microbe/views"
)

func (s *Server) setupRoutes() {
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.Logger)

	s.router.Get("/", docsHandler(views.IndexView))
	s.router.Get("/accent-color", docsHandler(views.AccentColorView))
	s.router.Get("/accordion", docsHandler(views.AccordionView))
	s.router.Get("/anchor", docsHandler(views.AnchorView))
	s.router.Get("/button", docsHandler(views.ButtonView))
	s.router.Get("/card", docsHandler(views.CardView))
	s.router.Get("/checkbox", docsHandler(views.CheckboxView))
	s.router.Get("/colors", docsHandler(views.Colors))
	s.router.Get("/dialog", docsHandler(views.DialogView))
	s.router.Get("/input", docsHandler(views.InputView))
	s.router.Get("/list", docsHandler(views.ListView))
	s.router.Get("/navlist", docsHandler(views.NavListView))
	s.router.Get("/progress", docsHandler(views.ProgressView))
	s.router.Get("/radio", docsHandler(views.RadioView))
	s.router.Get("/range", docsHandler(views.RangeView))
	s.router.Get("/select", docsHandler(views.SelectView))
	s.router.Get("/spacing", docsHandler(views.Spacing))
	s.router.Get("/switch", docsHandler(views.SwitchView))
	s.router.Get("/table", docsHandler(views.TableView))
	s.router.Get("/textarea", docsHandler(views.TextareaView))
	s.router.Get("/typography", docsHandler(views.Typography))

	s.router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
