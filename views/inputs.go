package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Inputs() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Inputs")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Example",
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
							),
							Small(Text("We’ll never share your email with anyone else.")),
						),
					),
					Div(
						Class("actions"),
						Button(Class("outline"), Type("reset"), Text("Reset")),
						Button(Class("solid"), Type("submit"), Text("Submit")),
					),
				),
			),

			example(
				"Types",
				Class("rows"),
				Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text")),
				Input(Type("email"), Name("email"), Aria("label", "Email"), Placeholder("Email"), AutoComplete("email")),
				Input(Type("number"), Name("number"), Aria("label", "Number"), Placeholder("Number")),
				Input(Type("password"), Name("password"), Aria("label", "Password"), Placeholder("Password")),
				Input(Type("search"), Name("search"), Aria("label", "Search"), Placeholder("Search")),
				Input(Type("tel"), Name("tel"), Aria("label", "Tel"), Placeholder("Tel")),
				Input(Type("url"), Name("url"), Aria("label", "Url"), Placeholder("Url")),
				Input(Type("date"), Name("date"), Aria("label", "Date")),
				Input(Type("datetime-local"), Name("datetime-local"), Aria("label", "Datetime local")),
				Input(Type("month"), Name("month"), Aria("label", "Month")),
				Input(Type("week"), Name("week"), Aria("label", "Week")),
				Input(Type("time"), Name("time"), Aria("label", "Time")),
				Input(Type("color"), Name("color"), Aria("label", "Color"), Placeholder("Color")),
				Input(Type("file"), Class("solid"), Name("file"), Aria("label", "File"), Placeholder("File")),
			),
			example(
				"Disabled",
				Class("rows"),
				Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text"), Disabled()),
				Input(Type("email"), Name("email"), Aria("label", "Email"), Placeholder("Email"), AutoComplete("email"), Disabled()),
				Input(Type("number"), Name("number"), Aria("label", "Number"), Placeholder("Number"), Disabled()),
				Input(Type("password"), Name("password"), Aria("label", "Password"), Placeholder("Password"), Disabled()),
				Input(Type("search"), Name("search"), Aria("label", "Search"), Placeholder("Search"), Disabled()),
				Input(Type("tel"), Name("tel"), Aria("label", "Tel"), Placeholder("Tel"), Disabled()),
				Input(Type("url"), Name("url"), Aria("label", "Url"), Placeholder("Url"), Disabled()),
				Input(Type("date"), Name("date"), Aria("label", "Date"), Disabled()),
				Input(Type("datetime-local"), Name("datetime-local"), Aria("label", "Datetime local"), Disabled()),
				Input(Type("month"), Name("month"), Aria("label", "Month"), Disabled()),
				Input(Type("week"), Name("week"), Aria("label", "Week"), Disabled()),
				Input(Type("time"), Name("time"), Aria("label", "Time"), Disabled()),
				Input(Type("color"), Name("color"), Aria("label", "Color"), Placeholder("Color"), Disabled()),
				Input(Type("file"), Class("solid"), Name("file"), Aria("label", "File"), Placeholder("File"), Disabled()),
			),
			example(
				"Readonly",
				Class("rows"),
				Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text"), ReadOnly(), Value("Read-only value")),
				Input(Type("email"), Name("email"), Aria("label", "Email"), Placeholder("Email"), AutoComplete("email"), ReadOnly()),
				Input(Type("number"), Name("number"), Aria("label", "Number"), Placeholder("Number"), ReadOnly()),
				Input(Type("password"), Name("password"), Aria("label", "Password"), Placeholder("Password"), ReadOnly()),
				Input(Type("search"), Name("search"), Aria("label", "Search"), Placeholder("Search"), ReadOnly()),
				Input(Type("tel"), Name("tel"), Aria("label", "Tel"), Placeholder("Tel"), ReadOnly()),
				Input(Type("url"), Name("url"), Aria("label", "Url"), Placeholder("Url"), ReadOnly()),
				Input(Type("date"), Name("date"), Aria("label", "Date"), ReadOnly()),
				Input(Type("datetime-local"), Name("datetime-local"), Aria("label", "Datetime local"), ReadOnly()),
				Input(Type("month"), Name("month"), Aria("label", "Month"), ReadOnly()),
				Input(Type("week"), Name("week"), Aria("label", "Week"), ReadOnly()),
				Input(Type("time"), Name("time"), Aria("label", "Time"), ReadOnly()),
			),
			example(
				"Hint",
				Class("rows"),
				Div(
					Input(Type("text"), Name("Hint"), Aria("label", "Hint"), Placeholder("Email")),
					Small(Text("We’ll never share your email with anyone else.")),
				),
			),
			example(
				"Validation",
				Class("rows"),
				Div(
					Input(Type("text"), Name("valid"), Aria("label", "Valid"), Aria("invalid", "false"), Aria("describedby", "valid-hint"), Value("Valid")),
					Small(ID("valid-hint"), Text("Looks good!")),
				),
				Div(
					Input(Type("text"), Name("invalid"), Aria("label", "Invalid"), Aria("invalid", "true"), Aria("describedby", "invalid-hint"), Value("Invalid")),
					Small(ID("invalid-hint"), Text("Please provide a valid value!")),
				),
			),
		),
	})
}
