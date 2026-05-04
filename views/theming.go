package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Theming() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Theming")),
		),
		Div(
			Attr("role", "document"),
			Label(
				Text("Hue"),
				Input(
					Attr("type", "range"),
					Attr("value", "240"),
					Attr("min", "0"),
					Attr("max", "359"),
					Attr("id", "color-range"),
				)),
			Label(
				Text("Saturation"),
				Input(
					Attr("type", "range"),
					Attr("value", "70%"),
					Attr("min", "0"),
					Attr("max", "100"),
					Attr("id", "saturation-range"),
				)),
			example(
				"Example",
				Attr("id", "theming-example-form"),
				Class("grid"),
				Button(Class("solid"), Text("Solid Button")),
				Button(Class("outline"), Text("Outline Button")),
				Button(Class("ghost"), Text("Ghost Button")),
			),
		),
	})
}
