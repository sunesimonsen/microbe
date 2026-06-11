package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AccordionView() Node {
	panelText := "Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex. Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce sagittis in est a consequat."

	return docpage(
		HGroup(
			H1(Text("Accordion")),
			P(Text("An element that organizes content into a vertically stacked list of collapsible sections. Users can click or tap a section's header to expand it and reveal detailed information, or collapse it to hide the content and reduce scrolling.")),
		),
		example(
			"Default",
			Section(
				Class("accordion"),
				Details(
					Summary(Text("Section 1")),
					P(Text(panelText)),
				),
				Details(
					Summary(Text("Section 2")),
					P(Text(panelText)),
				),
				Details(
					Summary(Text("Section 3")),
					P(Text(panelText)),
				),
				Details(
					Summary(Text("Section 4")),
					P(Text(panelText)),
				),
			),
		),
		example(
			"Single panel",
			Section(
				Class("accordion"),
				Details(
					Name("single"),
					Summary(Text("Section 1")),
					P(Text(panelText)),
				),
				Details(
					Name("single"),
					Summary(Text("Section 2")),
					P(Text(panelText)),
				),
				Details(
					Name("single"),
					Summary(Text("Section 3")),
					P(Text(panelText)),
				),
				Details(
					Name("single"),
					Summary(Text("Section 4")),
					P(Text(panelText)),
				),
			),
		),
	)
}
