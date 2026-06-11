package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CardView() Node {
	cardText := "Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat."

	return docpage(
		HGroup(
			H1(Text("Card")),
			P(Text("An element that groups related information and actions about a single subject into a visually distinct, flexible container.")),
		),
		example(
			"Bare card",
			Article(
				Class("card"),
				Section(P(Text(cardText))),
			),
		),
		example(
			"Card with header",
			Article(
				Class("card"),
				Header(Text("Header")),
				Section(P(Text(cardText))),
			),
		),
		example(
			"Card with footer",
			Article(
				Class("card"),
				Section(P(Text(cardText))),
				Footer(Text("Footer")),
			),
		),
		example(
			"Card with header and footer",
			Article(
				Class("card"),
				Header(Text("Header")),
				Section(P(Text(cardText))),
				Footer(Text("Footer")),
			),
		),
	)
}
