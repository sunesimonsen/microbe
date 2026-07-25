package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sunesimonsen/microbe/docs"
	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/getting-started/about", http.StatusMovedPermanently)
}

func writeInternalServerError(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func DocsHandler(w http.ResponseWriter, r *http.Request) {
	category, err := docs.Index.FindCategory(chi.URLParam(r, "category"))
	page, err := category.FindPage(chi.URLParam(r, "page"))

	if err != nil {
		if errors.Is(err, docs.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			writeInternalServerError(w)
		}
		return
	}

	renderNode(w, r, views.DocsLayout(r.URL.Path, page.GetNode(*r.URL)))
}

func renderNode(w http.ResponseWriter, _ *http.Request, node Node) {
	if node == nil {
		panic("renderNode without a node")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	if err := node.Render(w); err != nil {
		writeInternalServerError(w)
	}
}
