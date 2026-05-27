package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AnchorView() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Anchor")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Regular",
				A(Href("#"), Text("Regular link")),
			),
			example(
				"Active link",
				A(Href("#"), Attr("aria-current", "page"), Text("Active link")),
			),
		),
	})
}
