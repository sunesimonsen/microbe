package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AnchorView() Node {
	return docpage(
		HGroup(H1(Text("Anchor"))),
		NewExample(
			"Regular",
			A(Href("#"), Text("Regular link")),
		),
		NewExample(
			"Active link",
			A(Href("#"), Attr("aria-current", "page"), Text("Active link")),
		),
	)
}
