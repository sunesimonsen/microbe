package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func RangeView() Node {
	return docpage(
		HGroup(H1(Text("Range"))),
		NewExample(
			"Example",
			Nodes(
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
		),
		NewExample(
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
	)
}
