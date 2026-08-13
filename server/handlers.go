package server

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/sunesimonsen/microbe/docs"
	"github.com/sunesimonsen/microbe/views"
	. "maragu.dev/gomponents"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/docs/about", http.StatusMovedPermanently)
}

func writeInternalServerError(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func DocsHandler(w http.ResponseWriter, r *http.Request) {
	page, err := docs.Index.FindPage(r.PathValue("page"))

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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.Referer())

	path := ""

	if err == nil {
		path = u.Path
	}

	query := r.URL.Query().Get("query")
	expandAll := query != ""

	cs := docs.Index.Filter(query)

	if len(cs) == 0 {
		renderNode(w, r, views.NoSearchResults(query))
		return
	}

	menu := cs.GetMenu(path, expandAll)

	renderNode(w, r, views.SearchResults(menu))
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
