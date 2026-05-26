package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Ranges() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Ranges")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Example",
				Label(
					Text("Saturation"),
					Input(
						Attr("type", "range"),
						Attr("value", "70"),
						Attr("min", "0"),
						Attr("max", "100"),
						Aria("describedby", "saturation-hint"),
					)),
				Small(
					ID("saturation-hint"),
					Text("Accent color Saturation"),
				),
			),
			example(
				"Disabled",
				Label(
					Text("Saturation"),
					Input(
						Attr("type", "range"),
						Attr("value", "70"),
						Attr("min", "0"),
						Attr("max", "100"),
						Disabled(),
					)),
			),
		),
	})
}
