package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/sunesimonsen/microbe/templates"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Index()
	renderComponent(w, r, templates.IndexLayout(component))
}

func (s *Server) buttonsHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Buttons()
	renderComponent(w, r, templates.DocsLayout(component))
}

func (s *Server) headingsHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Headings()
	renderComponent(w, r, templates.DocsLayout(component))
}

func (s *Server) colorsHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Colors()
	renderComponent(w, r, templates.DocsLayout(component))
}

func (s *Server) spacingHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Spacing()
	renderComponent(w, r, templates.DocsLayout(component))
}

func (s *Server) anchorsHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Anchors()
	renderComponent(w, r, templates.DocsLayout(component))
}

func (s *Server) cardsHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Cards()
	renderComponent(w, r, templates.DocsLayout(component))
}

func renderComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	if component == nil {
		panic("renderComponent without a component")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
