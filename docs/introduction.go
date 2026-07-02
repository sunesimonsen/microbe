package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var IntroductionPage = NewPage(
	"Introduction",
	nil,
	NewPageSection(
		"Introduction",
		"Install",
		Div(
			Role("document"),
			H2(Text("!!! Under construction !!!")),
			Pre(Code(Text("https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css"))),
		),
	),
)
