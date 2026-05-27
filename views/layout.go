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
				Li(A(Href("/button"), Text("Button"))),
				Li(A(Href("/anchor"), Text("Anchor"))),
				Li(A(Href("/input"), Text("Input"))),
				Li(A(Href("/checkbox"), Text("Checkbox"))),
				Li(A(Href("/switch"), Text("Switch"))),
				Li(A(Href("/radio"), Text("Radio"))),
				Li(A(Href("/select"), Text("Select"))),
				Li(A(Href("/range"), Text("Range"))),
				Li(A(Href("/spacing"), Text("Spacing"))),
				Li(A(Href("/card"), Text("Card"))),
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
