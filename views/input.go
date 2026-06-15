package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func InputView() Node {
	return docpage(
		HGroup(H1(Text("Input"))),
		example(
			"Form",
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
				),
				Div(
					Class("actions"),
					Button(Class("outline"), Type("reset"), Text("Reset")),
					Button(Class("solid"), Type("submit"), Text("Submit")),
				),
			),
		),

		example(
			"Text inputs",
			Class("rows"),
			Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text")),
			Input(Type("email"), Name("email"), Aria("label", "Email"), Placeholder("Email"), AutoComplete("email")),
			Input(Type("number"), Name("number"), Aria("label", "Number"), Placeholder("Number")),
			Input(Type("password"), Name("password"), Aria("label", "Password"), Placeholder("Password")),
			Input(Type("tel"), Name("tel"), Aria("label", "Tel"), Placeholder("Tel")),
			Input(Type("url"), Name("url"), Aria("label", "Url"), Placeholder("Url")),
		),

		example(
			"Date and time inputs",
			Class("rows"),
			Input(Type("date"), Name("date"), Aria("label", "Date")),
			Input(Type("datetime-local"), Name("datetime-local"), Aria("label", "Datetime local")),
			Input(Type("month"), Name("month"), Aria("label", "Month")),
			Input(Type("week"), Name("week"), Aria("label", "Week")),
			Input(Type("time"), Name("time"), Aria("label", "Time")),
		),

		example(
			"Search input",
			Input(Type("search"), Name("search"), Aria("label", "Search"), Placeholder("Search")),
		),

		example(
			"File input",
			Input(Type("file"), Class("solid"), Name("file"), Aria("label", "File"), Placeholder("File")),
		),

		example(
			"Color input",
			Input(Type("color"), Name("color"), Aria("label", "Color"), Placeholder("Color")),
		),

		example(
			"Disabled input",
			Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text"), Disabled()),
		),
		example(
			"Read-only",
			Input(Type("text"), Name("text"), Aria("label", "Text"), Placeholder("Text"), ReadOnly(), Value("Read-only value")),
		),
		example(
			"Hint",
			Input(Type("text"), Name("Hint"), Aria("label", "Hint"), Placeholder("Email")),
			Small(Text("We’ll never share your email with anyone else.")),
		),
		example(
			"Validation",
			Input(Type("text"), Name("valid"), Aria("label", "Valid"), Aria("invalid", "false"), Aria("describedby", "valid-hint"), Value("Valid")),
			Small(ID("valid-hint"), Text("Looks good!")),
			Input(Type("text"), Name("invalid"), Aria("label", "Invalid"), Aria("invalid", "true"), Aria("describedby", "invalid-hint"), Value("Invalid")),
			Small(ID("invalid-hint"), Text("Please provide a valid value!")),
		),
	)
}
