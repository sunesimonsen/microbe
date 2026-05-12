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
				Li(A(Href("/typography"), Text("Typography"))),
				Li(A(Href("/buttons"), Text("Buttons"))),
				Li(A(Href("/anchors"), Text("Anchors"))),
				Li(A(Href("/inputs"), Text("Inputs"))),
				Li(A(Href("/checkboxes"), Text("Checkboxes"))),
				Li(A(Href("/spacing"), Text("Spacing"))),
				Li(A(Href("/cards"), Text("Cards"))),
				Li(A(Href("/colors"), Text("Colors"))),
				Li(A(Href("/theming"), Text("Theming"))),
			),
		),
	)
}

func example(description string, part ...Node) Node {
	return Article(
		Class("example card"),
		Header(Text(description)),
		Section(
			part...,
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
