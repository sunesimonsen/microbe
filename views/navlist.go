package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func NavListView() Node {
	return docpage(
		HGroup(H1(Text("Navlist"))),
		example(
			"Example",
			Nav(
				Class("navlist"),
				Details(
					Summary(Text("Content")),
					Ul(
						Li(A(Href("/typography"), Text("Typography"))),
					),
				),
				Details(
					Open(),
					Summary(Text("Navigation")),
					Ul(
						Li(A(Href("/anchor"), Text("Anchor"))),
						Li(A(Href("/navlist"), Aria("current", "page"), Text("Navlist"))),
					),
				),
				Details(
					Summary(Text("Layout")),
					Ul(
						Li(A(Href("/spacing"), Text("Spacing"))),
					),
				),
				Details(
					Summary(Text("Forms")),
					Ul(
						Li(A(Href("/button"), Text("Button"))),
						Li(A(Href("/checkbox"), Text("Checkbox"))),
						Li(A(Href("/input"), Text("Input"))),
						Li(A(Href("/radio"), Text("Radio"))),
						Li(A(Href("/range"), Text("Range"))),
						Li(A(Href("/select"), Text("Select"))),
						Li(A(Href("/switch"), Text("Switch"))),
					),
				),
				Details(
					Summary(Text("Components")),
					Ul(
						Li(A(Href("/card"), Text("Card"))),
						Li(A(Href("/dialog"), Text("Dialog"))),
					),
				),
				Details(
					Summary(Text("Theming")),
					Ul(
						Li(A(Href("/accent-color"), Text("Accent color"))),
						Li(A(Href("/colors"), Text("Colors"))),
					),
				),
			),
		),
	)
}
