package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IndexView() Node {
	return Nodes(
		HGroup(
			H1(Text("Microbe")),
		),
		Div(
			Role("document"),
			H2(Text("Hello")),
		),
	)
}
