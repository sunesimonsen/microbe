package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Anchors() Node {
	return Group([]Node{
		HGroup(
			ID("content-header"),
			H1(Text("Anchors")),
		),
		Div(
			ID("content"),
			Attr("role", "document"),
			Div(
				Class("anchors-row"),
				A(Href("#"), Text("Regular link")),
				A(Href("#"), Attr("aria-current", "page"), Text("Active link")),
			),
		),
	})
}
