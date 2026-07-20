package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var IntroductionPage = NewPage(
	"Introduction",
	NewPageSection(
		"Usage",
		Div(
			Role("document"),
			H2(Text("Usage")),
			P(Text("This project is currently under construction, so no releases has been made yet. But you can find the development stylesheets here if you want to play around with it")),
			P(Text("The main stylesheet is "), Code(Text("microbe.css")),
				Text(" is required and provides base styles for most HTML elements.")),
			Pre(Code(Text("https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css"))),
			P(Text("A number of additional stylesheets exists extending the main stylesheet with extra components. You can include these to meet your needs:")),
			Ul(
				Li(A(Href("/components/accordion"), Text("Accordion"))),
				Li(A(Href("/components/card"), Text("Card"))),
				Li(A(Href("/navigation/navlist"), Text("Navlist"))),
			),
		),
	),
)
