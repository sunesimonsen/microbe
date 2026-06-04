package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AnchorView() Node {
	return docpage(
		HGroup(H1(Text("Anchor"))),
		example(
			"Regular",
			A(Href("#"), Text("Regular link")),
		),
		example(
			"Active link",
			A(Href("#"), Attr("aria-current", "page"), Text("Active link")),
		),
	)
}
