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
					Attr("value", "70"),
					Attr("min", "0"),
					Attr("max", "100"),
					Attr("id", "saturation-range"),
				)),
			example(
				"Example",
				Attr("id", "theming-example-form"),
				Form(
					FieldSet(
						Label(Text("Name"),
							Input(
								Name("name"),
								Placeholder("Name"),
								AutoComplete("name"),
							),
						),
						Label(Text("Email"),
							Input(
								Name("email"),
								Placeholder("Email"),
								AutoComplete("email"),
								Aria("describedby", "email-hint"),
							),
							Small(
								ID("email-hint"),
								Text("We’ll never share your email with anyone else."),
							),
						),
						Label(
							Input(
								Type("checkbox"),
								Name("newsletter"),
								Aria("describedby", "newsletter-hint"),
								Checked(),
							),
							Text("Newsletter"),
						),
						Small(
							ID("newsletter-hint"),
							Text("We will send you a news letter every week"),
						),
						Label(
							Text("Spelling proficiency (0 - 10)"),
							Input(
								Attr("type", "range"),
								Attr("value", "5"),
								Attr("min", "0"),
								Attr("max", "10"),
							)),
					),
					Div(
						Class("actions"),
						Button(Class("outline"), Type("reset"), Text("Reset")),
						Button(Class("solid"), Type("submit"), Text("Submit")),
					),
				),
			),
		),
	})
}
