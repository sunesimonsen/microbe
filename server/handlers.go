package server

import (
	"net/http"

	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

type view func() Node

func docsHandler(v view) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderNode(w, r, views.DocsLayout(r.URL.Path, v()))
	}
}

func renderNode(w http.ResponseWriter, _ *http.Request, node Node) {
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
