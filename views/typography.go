package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Typography() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Typography")),
			P(Text("Fill in some info here")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Headings",
				H1(Text("Heading 1")),
				P(Text("Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.")),
				H2(Text("Heading 2")),
				P(Text("Phasellus nec luctus dolor. Curabitur id facilisis diam.")),
				H3(Text("Heading 3")),
				P(Text("Vivamus feugiat tempor tellus, vel consequat turpis gravida non. ")),
				H4(Text("Heading 4")),
				P(Text("Quisque tristique lobortis ligula id tempus.")),
				H5(Text("Heading 5")),
				P(Text("Sed aliquet velit mauris, vel interdum diam mattis et.")),
				H6(Text("Heading 6")),
				P(Text("Donec in lorem imperdiet, eleifend turpis eget, congue velit.")),
			),
		),
	})
}
