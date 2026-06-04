package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ProgressView() Node {
	return docpage(
		HGroup(H1(Text("Progress"))),
		example(
			"Example",
			Label(
				Text("Downloading file"),
				Progress(Attr("value", "70"), Attr("max", "100")),
			),
		),
		example(
			"Indeterminate",
			Label(
				Text("Downloading file"),
				Progress(),
			),
		),
	)
}
