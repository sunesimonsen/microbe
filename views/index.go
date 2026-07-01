package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IndexView() Node {
	return docpage(
		HGroup(H1(Text("Introduction"))),
		NewPageSection("sdfsdf",
			Div(
				Role("document"),
				H2(Text("!!! Under construction !!!")),
				Pre(Code(Text("https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css"))),
			),
		),
	)
}
