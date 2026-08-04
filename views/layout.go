package views

import (
	"github.com/sunesimonsen/microbe/docs"
	"github.com/sunesimonsen/microbe/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func header() Node {
	return Header(
		Nav(
			Div(
				Class("header-items"),
				A(Class("ghost"), Title("Home"), Href("/"), Text("Microbe")),
			),
			Div(
				Class("header-items"),
				A(
					Href("https://github.com/sunesimonsen/microbe"),
					Target("_blank"),
					Title("Github"),
					Class("ghost"),
					Span(Class("only-wide"), Text("Github ")),
					icons.GithubIcon(),
				),
				Button(
					ID("search-button"),
					Class("ghost"),
					Title("Search"),
					Attr("command", "show-modal"),
					Attr("commandfor", "search-dialog"),
					Span(Class("only-wide"), Text("Search ")),
					Span(Class("only-wide"), icons.SearchIcon()),
					Span(Class("only-narrow"), icons.BurgerIcon()),
				),
			),
		),
		Dialog(
			ID("search-dialog"),
			Attr("closedby", "any"),
			Header(
				Label(
					Text("Search"),
					Input(
						Type("search"),
						Name("query"),
						AutoFocus(),
						Placeholder("Search for documentation"),
						Attr("autocomplete", "off"),
						Attr("hx-get", "/search"),
						Attr("hx-params", "query"),
						Attr("hx-trigger", "input changed delay:500ms, keyup[key=='Enter'], load"),
						Attr("hx-target", "#search-results"),
					),
				),
				Button(Rel("prev"), Aria("label", "Close"), Attr("commandfor", "search-dialog"), Attr("command", "close"), TabIndex("1")),
			),
			Section(
				ID("search-results"),
			),
		),
	)
}

func docsMenu(currentPath string) Node {
	return Aside(
		Class("menu"),

		Nav(
			Class("navlist"),
			docs.Index.GetMenu(currentPath, false),
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

func SearchResults(part Node) Node {
	return Nav(
		Class("navlist"),
		part,
	)
}

func NoSearchResults(query string) Node {
	return P(Class("no-results"), Text("No results for \""), Em(Text(query)), Text("\""))
}
