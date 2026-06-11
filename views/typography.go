package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Typography() Node {
	return docpage(
		HGroup(
			H1(Text("Typography")),
			P(Text("Fill in some info here")),
		),
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
		example(
			"Inline text",
			Class("grid"),
			Abbr(Text("Abbreviation")),
			Strong(Text("Strong")),
			Em(Text("Emphasised")),
			Del(Text("Deleted")),
			Ins(Text("Inserted")),
			Kbd(Text("Ctrl + S")),
			Mark(Text("Highlighted")),
			S(Text("Strikethrough")),
			U(Text("Underlined")),
			Span(Text("X"), Sub(Text("sub"))),
			Span(Text("X"), Sup(Text("sup"))),
			Small(Text("Small")),
			Samp(Text("Samp")),
			Code(Text("Code")),
			Var(Text("Var")),
		),
	)
}
