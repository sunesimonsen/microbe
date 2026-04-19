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
				Li(A(Href("/headings"), Attr("data-swap", "content-header,content"), Text("Headings"))),
				Li(A(Href("/buttons"), Attr("data-swap", "content-header,content"), Text("Buttons"))),
				Li(A(Href("/anchors"), Attr("data-swap", "content-header,content"), Text("Anchors"))),
				Li(A(Href("/spacing"), Attr("data-swap", "content-header,content"), Text("Spacing"))),
				Li(A(Href("/colors"), Attr("data-swap", "content-header,content"), Text("Colors"))),
				Li(A(Href("/cards"), Attr("data-swap", "content-header,content"), Text("Cards"))),
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
