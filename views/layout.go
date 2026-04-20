package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func header() Node {
	return Header(
		Nav(
			Button(
				Class("menu-toggle ghost icon"),
				burgerIcon(),
			),
			Text("Microbe"),
		),
	)
}

func menu() Node {
	return Aside(
		Class("menu"),
		Nav(
			Ul(
				Li(A(Href("/headings"), Text("Headings"))),
				Li(A(Href("/buttons"), Text("Buttons"))),
				Li(A(Href("/anchors"), Text("Anchors"))),
				Li(A(Href("/spacing"), Text("Spacing"))),
				Li(A(Href("/colors"), Text("Colors"))),
				Li(A(Href("/cards"), Text("Cards"))),
			),
		),
	)
}

func IndexLayout(part Node) Node {
	return Page("Microbe",
		Main(
			Class("standard-layout"),
			header(),
			part,
		),
	)
}

func DocsLayout(part Node) Node {
	return Page("Microbe",
		Main(
			Class("standard-layout"),
			header(),
			menu(),
			part,
		),
	)
}
