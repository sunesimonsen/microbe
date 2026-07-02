package views

import (
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/sunesimonsen/microbe/docs"
	"github.com/sunesimonsen/microbe/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func header() Node {
	return Header(
		Nav(
			Button(
				Class("menu-toggle ghost icon"),
				icons.BurgerIcon(),
			),
			A(Class("home"), Href("/"), Text("Microbe")),
		),
	)
}

func PageHref(c docs.Category, p docs.Page) string {
	return "/" + strcase.ToKebab(c.Name) + "/" + strcase.ToKebab(p.Name)
}

func docsMenu(currentPath string) Node {
	return Aside(
		Class("menu"),

		Nav(
			Class("navlist"),
			Map(docs.Index, func(c docs.Category) Node {
				hasActivePage := slices.ContainsFunc(c.Pages, func(p docs.Page) bool {
					return PageHref(c, p) == currentPath
				})

				return Details(
					Name("index"),
					If(hasActivePage, Open()),
					Summary(Text(c.Name)),
					Ul(
						Map(c.Pages, func(p docs.Page) Node {
							href := "/" + strcase.ToKebab(c.Name) + "/" + strcase.ToKebab(p.Name)
							return Li(A(
								If(PageHref(c, p) == currentPath, Aria("current", "page")),
								Href(href),
								Text(p.Name),
							))
						}),
					),
				)
			}),
		),
	)
}

func DocsLayout(currentPath string, part Node) Node {
	return Page("Microbe",
		Main(
			Class("page-layout"),
			header(),
			docsMenu(currentPath),
			part,
		),
	)
}
