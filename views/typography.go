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
		NewExample(
			"Headings",
			Nodes(
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
		NewExample(
			"Inline text",
			Nodes(
				Class("grid"),
				Abbr(Text("Abbreviation")),
				Cite(Text("Citation")),
				Code(Text("Code")),
				Del(Text("Deleted")),
				Dfn(Text("Definition")),
				Em(Text("Emphasised")),
				I(Text("Idiomatic")),
				Ins(Text("Inserted")),
				Kbd(Text("Ctrl + S")),
				Mark(Text("Highlighted")),
				S(Text("Strikethrough")),
				Samp(Text("Samp")),
				Small(Text("Small")),
				Span(Text("X"), Sub(Text("sub"))),
				Span(Text("X"), Sup(Text("sup"))),
				Strong(Text("Strong")),
				U(Text("Underlined")),
				Var(Text("Var")),
			),
		),
		NewExample(
			"Horizontal ruler",
			Nodes(
				P(Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla ullamcorper blandit ultricies. Etiam non suscipit felis. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Proin feugiat purus hendrerit sapien condimentum, eu facilisis ex eleifend. Sed lobortis est a urna accumsan vestibulum. Curabitur iaculis sem lacus, id molestie eros ultrices eu. Nullam justo nulla, sollicitudin sed mi pretium, consequat varius erat. Phasellus sed eros dictum, congue enim in, dictum ante. Sed condimentum ut elit eu vehicula.")),
				Hr(),
				P(Text("Fusce congue nec massa id eleifend. Donec id luctus ligula. In pellentesque diam eu magna interdum, a tristique arcu dignissim. In eu lacinia nisl. Phasellus eu nisi vitae enim aliquam rutrum. Suspendisse fringilla tortor et tincidunt vulputate. Nam a urna a purus ornare gravida non sit amet risus. Morbi in justo quis velit elementum fermentum. Etiam tristique diam nunc, quis suscipit velit suscipit non.")),
			),
		),
		NewExample(
			"Heading group",
			HGroup(
				H2(Text("Get inspired with CSS")),
				P(Text("How to use CSS to add glam to your Website?")),
			),
		),
		NewExample(
			"Blockquote",
			BlockQuote(
				Text("\"Be the change that you wish to see in the world.\""),
				Footer(Cite(Text("— Mahatma Gandhi"))),
			),
		),
		NewExample(
			"Code block",
			Pre(Code(Text("console.log('Hello world!');"))),
		),
	)
}
