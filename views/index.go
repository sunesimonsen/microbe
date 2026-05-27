package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IndexView() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Microbe")),
		),
		Div(
			Attr("role", "document"),
			H2(Text("Hello")),
		),
	})
}
