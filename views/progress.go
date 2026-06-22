package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ProgressView() Node {
	return docpage(
		HGroup(H1(Text("Progress"))),
		NewExample(
			"Example",
			Label(
				Text("Downloading file"),
				Progress(Attr("value", "70"), Attr("max", "100")),
			),
		),
		NewExample(
			"Indeterminate",
			Label(
				Text("Downloading file"),
				Progress(),
			),
		),
	)
}
