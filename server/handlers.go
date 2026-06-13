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
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.ButtonView()))
}

func (s *Server) typographyHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.Typography()))
}

func (s *Server) tableHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.TableView()))
}

func (s *Server) colorsHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.Colors()))
}

func (s *Server) spacingHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.Spacing()))
}

func (s *Server) anchorHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.AnchorView()))
}

func (s *Server) cardHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.CardView()))
}

func (s *Server) accentColorHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.AccentColorView()))
}

func (s *Server) inputHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.InputView()))
}

func (s *Server) checkboxHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.CheckboxView()))
}

func (s *Server) switchHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.SwitchView()))
}

func (s *Server) radioHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.RadioView()))
}

func (s *Server) selectHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.SelectView()))
}

func (s *Server) rangeHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.RangeView()))
}

func (s *Server) progressHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.ProgressView()))
}

func (s *Server) dialogHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.DialogView()))
}

func (s *Server) navlistHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.NavListView()))
}

func (s *Server) accordionHandler(w http.ResponseWriter, r *http.Request) {
	renderNode(w, r, views.DocsLayout(r.URL.Path, views.AccordionView()))
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
