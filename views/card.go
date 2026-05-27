package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CardView() Node {
	cardText := "Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat."

	return Group([]Node{
		HGroup(
			H1(Text("Card")),
		),
		Div(
			Attr("role", "document"),
			Article(
				Class("card"),
				Section(
					H5(Text("Example card")),
					P(Text(cardText)),
				),
			),
			Article(
				Class("card"),
				Header(Text("Header")),
				Section(
					H5(Text("Example card with header")),
					P(Text(cardText)),
				),
			),
			Article(
				Class("card"),
				Section(
					H5(Text("Example card with footer")),
					P(Text(cardText)),
				),
				Footer(Text("Footer")),
			),
			Article(
				Class("card"),
				Header(Text("Header")),
				Section(
					H5(Text("Example card with header and footer")),
					P(Text(cardText)),
				),
				Footer(Text("Footer")),
			),
		),
	})
}
