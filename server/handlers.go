package server

import (
	"net/http"

	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.IndexLayout(views.Index()))
}

func (s *Server) buttonsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Buttons()))
}

func (s *Server) typographyHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Typography()))
}

func (s *Server) colorsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Colors()))
}

func (s *Server) spacingHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Spacing()))
}

func (s *Server) anchorsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Anchors()))
}

func (s *Server) cardsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Cards()))
}

func (s *Server) themingHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Theming()))
}

func (s *Server) inputsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Inputs()))
}

func (s *Server) checkboxesHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Checkboxes()))
}

func (s *Server) radiosHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Radios()))
}

func (s *Server) selectsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Selects()))
}

func renderNode(w http.ResponseWriter, r *http.Request, node Node) {
	if node == nil {
		panic("renderNode without a node")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	if err := node.Render(w); err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
