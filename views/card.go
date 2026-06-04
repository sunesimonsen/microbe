package views

import (
	"github.com/iancoleman/strcase"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func exampleCard(name string, parts ...Node) pageSection {
	return pageSection{
		name: name,
		content: Article(
			ID(strcase.ToKebab(name)),
			Class("card"),
			Group(parts),
		),
	}
}

func CardView() Node {
	cardText := "Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat."

	return docpage(
		HGroup(H1(Text("Card"))),
		exampleCard(
			"bare card",
			Section(
				H5(Text("Bare card")),
				P(Text(cardText)),
			),
		),
		exampleCard(
			"Card with header",
			Header(Text("Header")),
			Section(
				H5(Text("Card with header")),
				P(Text(cardText)),
			),
		),
		exampleCard(
			"Card with footer",
			Section(
				H5(Text("Card with footer")),
				P(Text(cardText)),
			),
			Footer(Text("Footer")),
		),
		exampleCard(
			"Card with header and footer",
			Header(Text("Header")),
			Section(
				H5(Text("Card with header and footer")),
				P(Text(cardText)),
			),
			Footer(Text("Footer")),
		),
	)
}
