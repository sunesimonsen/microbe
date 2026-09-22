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
				A(Class("home"), Title("Home"), Href("/"), Img(Src("/assets/microbe-cube.svg"), Alt("Logo")), Text("Microbe")),
			),
			Div(
				Class("header-items"),
				A(
					Href("https://github.com/sunesimonsen/microbe"),
					Target("_blank"),
					Class("ghost icon"),
					Title("Github"),
					icons.GithubIcon(),
				),
				Button(
					ID("theme-button"),
					Attr("popovertarget", "theme-menu"),
					Attr("aria-controls", "theme-menu"),
					Attr("aria-haspopup", "menu"),
					Class("ghost icon"),
					Span(Class("theme-icon theme-icon-system"), Attr("aria-hidden", "true"), icons.CircleHalfFillIcon()),
					Span(Class("theme-icon theme-icon-light"), Attr("aria-hidden", "true"), icons.SunIcon()),
					Span(Class("theme-icon theme-icon-dark"), Attr("aria-hidden", "true"), icons.MoonIcon()),
				),
				Ul(
					ID("theme-menu"),
					Class("menu"),
					Attr("popover", "auto"),
					Attr("role", "menu"),
					Attr("aria-label", "Color theme"),
					Li(Button(Type("button"), Attr("data-theme", "system"), Attr("role", "menuitemradio"), Aria("checked", "true"),
						Span(Attr("aria-hidden", "true"), icons.CircleHalfFillIcon()), Text("System"))),
					Li(Button(Type("button"), Attr("data-theme", "light"), Attr("role", "menuitemradio"), Aria("checked", "false"),
						Span(Attr("aria-hidden", "true"), icons.SunIcon()), Text("Light"))),
					Li(Button(Type("button"), Attr("data-theme", "dark"), Attr("role", "menuitemradio"), Aria("checked", "false"),
						Span(Attr("aria-hidden", "true"), icons.MoonIcon()), Text("Dark"))),
				),
				Button(
					ID("search-button"),
					Attr("command", "show-modal"),
					Attr("commandfor", "search-dialog"),
					Attr("accesskey", "s"),
					Class("ghost icon"),
					Title("Search"),
					icons.BurgerIcon(),
				),
			),
		),
		Dialog(
			ID("search-dialog"),
			Class("search"),
			Attr("closedby", "any"),
			Header(
				P(Label(For("search-input"), Text("Search"))),
				Input(
					ID("search-input"),
					Type("search"),
					Name("query"),
					AutoFocus(),
					Placeholder("Search for documentation"),
					Attr("autocomplete", "off"),
					Attr("hx-get", "/search"),
					Attr("hx-params", "query"),
					Attr("hx-trigger", "input changed delay:500ms, keyup[key=='Enter'], intersect once"),
					Attr("hx-target", "#search-results"),
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
		header(),
		Main(
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
