package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ColorsPage = NewPage(
	"Colors",
	NewStaticPageSection(
		"Playground",
		Article(
			H2(Text("Playground")),
			Div(
				Class("card raised"),
				Header(
					Label(
						Text("Accent hue"),
						Input(
							Attr("type", "range"),
							Attr("value", "240"),
							Attr("min", "0"),
							Attr("max", "359"),
							Attr("id", "accent-color-range"),
							Class("color-range"),
						)),
					Label(
						Text("Accent saturation"),
						Input(
							Attr("type", "range"),
							Attr("value", "70"),
							Attr("min", "0"),
							Attr("max", "100"),
							Attr("id", "accent-saturation-range"),
							Class("saturation-range"),
						)),
					Hr(),
					Label(
						Text("Neutral hue"),
						Input(
							Attr("type", "range"),
							Attr("value", "0"),
							Attr("min", "0"),
							Attr("max", "359"),
							Attr("id", "neutral-color-range"),
							Style("--accent-hue: 0; --accent-saturation: 0"),
							Class("color-range"),
						)),
					Label(
						Text("Neutral saturation"),
						Input(
							Attr("type", "range"),
							Attr("value", "0"),
							Attr("min", "0"),
							Attr("max", "100"),
							Attr("id", "neutral-saturation-range"),
							Class("saturation-range"),
						)),
				),
				Section(
					Form(
						ID("theming-example-form"),
						FieldSet(
							Label(
								Text("Name"),
								Input(Placeholder("Name"), AutoComplete("name")),
							),
							Label(
								Text("Email"),
								Input(Placeholder("Email"), AutoComplete("email"), Aria("describedby", "email-hint")),
								Small(ID("email-hint"), Text("We'll never share your email with anyone else.")),
							),
							Label(
								Input(Type("checkbox"), Checked(), Aria("describedby", "newsletter-hint")),
								Text("Newsletter"),
							),
							Small(ID("newsletter-hint"), Text("We will send you a newsletter every week.")),
						),
						Div(
							Class("actions"),
							Button(Class("outline"), Type("reset"), Text("Reset")),
							Button(Class("solid"), Type("submit"), Text("Submit")),
						),
					),
				),
			),
		),
	),
).WithDescription(
	P(Text("Interactively adjust the accent and neutral hue and saturation to see how the whole color theme updates live.")),
)
