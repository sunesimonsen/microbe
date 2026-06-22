package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TextareaView() Node {
	return docpage(
		HGroup(H1(Text("Textarea"))),
		NewExample(
			"Default",
			Label(
				Text("Tell us your story:"),
				Textarea(Placeholder("It was a dark and stormy night...")),
			),
		),
		NewExample(
			"Custom rows",
			Label(
				Text("Tell us your story:"),
				Textarea(Rows("7"), Placeholder("It was a dark and stormy night...")),
			),
		),
		NewExample(
			"Disabled",
			Label(
				Text("Tell us your story:"),
				Textarea(Disabled(), Placeholder("It was a dark and stormy night...")),
			),
		),
		NewExample(
			"Read-only",
			Label(
				Text("Tell us your story:"),
				Textarea(ReadOnly(), Placeholder("It was a dark and stormy night..."), Text("Read-only value")),
			),
		),
		NewExample(
			"Hint",
			Nodes(
				Label(
					Text("Tell us your story:"),
					Textarea(Placeholder("It was a dark and stormy night...")),
				),
				Small(Text("Scary stories are often more engaging")),
			),
		),
		NewExample(
			"Validation",
			Nodes(
				Textarea(Type("text"), Name("valid"), Aria("label", "Valid"), Aria("invalid", "false"), Aria("describedby", "valid-hint"), Text("Valid")),
				Small(ID("valid-hint"), Text("Looks good!")),
				Textarea(Type("text"), Name("invalid"), Aria("label", "Invalid"), Aria("invalid", "true"), Aria("describedby", "invalid-hint"), Text("Invalid")),
				Small(ID("invalid-hint"), Text("Please provide a valid value!")),
			),
		),
	)
}
