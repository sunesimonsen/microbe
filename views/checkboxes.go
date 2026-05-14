package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Checkboxes() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Checkboxes")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Example",
				FieldSet(
					Legend(Text("Language preferences:")),
					Label(
						Input(
							Type("checkbox"),
							Name("english"),
							Checked(),
							Disabled(),
							Aria("describedby", "default-language-hint"),
						),
						Text("English"),
					),
					Small(
						ID("default-language-hint"),
						Text("You can't disable the default language"),
					),
					Label(
						Input(
							Type("checkbox"),
							Name("french"),
							Checked(),
						),
						Text("French"),
					),
					Label(
						Input(
							Type("checkbox"),
							Name("mandarin"),
						),
						Text("Mandarin"),
					),
					Label(
						Input(
							Type("checkbox"),
							Name("thai"),
						),
						Text("Thai"),
					),
					Label(
						Input(
							Type("checkbox"),
							Name("dothraki"),
							Disabled(),
						),
						Text("Dothraki"),
					),
				),
			),
			example(
				"Hint",
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
			example(
				"Validation",
				Label(
					Input(
						Type("checkbox"),
						Name("valid"),
						Aria("invalid", "false"),
						Aria("describedby", "valid-hint"),
					),
					Text("Valid"),
				),
				Small(ID("valid-hint"), Text("Looks good!")),
				Label(
					Input(
						Type("checkbox"),
						Name("invalid"),
						Aria("invalid", "true"),
						Aria("describedby", "invalid-hint"),
						Checked(),
					),
					Text("Invalid"),
				),
				Small(ID("invalid-hint"), Text("Please provide a valid value!")),
			),
		),
	})
}
