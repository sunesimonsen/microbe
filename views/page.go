package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Page(title string, children ...Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Class("microbe"),
			Head(
				TitleEl(Text(title)),
				Meta(Charset("UTF-8")),
				Meta(Name("mobile-web-app-capable"), Content("yes")),
				Meta(Name("viewport"), Content("width=device-width,initial-scale=1")),
				Meta(Name("description"), Content("Microbe.css documentation"), Lang("en")),
				Meta(Name("color-scheme"), Content("light dark")),
				Link(Rel("icon"), Type("image/png"), Href("/assets/favicon-96x96.png"), Attr("sizes", "96x96")),
				Link(Rel("icon"), Type("image/svg+xml"), Href("/assets/favicon.svg")),
				Link(Rel("shortcut icon"), Href("/assets/favicon.ico")),
				Link(Rel("apple-touch-icon"), Attr("sizes", "180x180"), Href("/assets/apple-touch-icon.png")),
				Meta(Name("apple-mobile-web-app-title"), Content("Microbe")),
				Link(Rel("manifest"), Href("/assets/site.webmanifest")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-accordion.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-breadcrumb.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-card.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-dialog.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-navlist.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-pagination.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-tabs.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-tag.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/microbe-menu.css")),
				Link(Rel("stylesheet"), Type("text/css"), Href("/assets/page.css")),
			),
			Body(
				Group(children),
				Script(Src("/assets/page.js")),
				Script(
					Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js"),
					Integrity("sha384-H5SrcfygHmAuTDZphMHqBJLc3FhssKjG7w/CeCpFReSfwBWDTKpkzPP8c+cLsK+V"),
					CrossOrigin("anonymous"),
				),
			),
		),
	)
}
