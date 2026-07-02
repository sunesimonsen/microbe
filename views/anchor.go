package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AnchorView() Node {
	return docpage(
		HGroup(H1(Text("Anchor"))),
		NewExample2(
			"anchor",
			"Regular",
		),
		NewExample2(
			"anchor",
			"Active link",
			WithDescription2(
				P(Text("Use "), Code(Text("aria-current=\"page\"")), Text(" to indicate anchor is pointing the current page.")),
			),
		),
	)
}
