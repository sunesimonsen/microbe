package server

import (
	"net/http"

	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.IndexLayout(views.IndexView()))
}

func (s *Server) buttonHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.ButtonView()))
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

func (s *Server) anchorHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.AnchorView()))
}

func (s *Server) cardHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.CardView()))
}

func (s *Server) themingHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.Theming()))
}

func (s *Server) inputHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.InputView()))
}

func (s *Server) checkboxHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.CheckboxView()))
}

func (s *Server) switchHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.SwitchView()))
}

func (s *Server) radioHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.RadioView()))
}

func (s *Server) selectHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.SelectView()))
}

func (s *Server) rangeHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.RangeView()))
}

func (s *Server) progressHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(views.ProgressView()))
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
