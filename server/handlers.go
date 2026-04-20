package server

import (
	"net/http"

	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Index()
	renderNode(w, r, views.IndexLayout(node))
}

func (s *Server) buttonsHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Buttons()
	renderNode(w, r, views.DocsLayout(node))
}

func (s *Server) typographyHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Typography()
	renderNode(w, r, views.DocsLayout(node))
}

func (s *Server) colorsHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Colors()
	renderNode(w, r, views.DocsLayout(node))
}

func (s *Server) spacingHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Spacing()
	renderNode(w, r, views.DocsLayout(node))
}

func (s *Server) anchorsHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Anchors()
	renderNode(w, r, views.DocsLayout(node))
}

func (s *Server) cardsHandler(w http.ResponseWriter, r *http.Request) {
	node := views.Cards()
	renderNode(w, r, views.DocsLayout(node))
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
