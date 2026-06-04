package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func RadioView() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Radio")),
		),
		Div(
			Role("document"),
			example(
				"Example",
				FieldSet(
					Legend(Text("Language preferences:")),
					Label(
						Input(
							Type("radio"),
							Name("language"),
							Checked(),
							Aria("describedby", "default-language-hint"),
						),
						Text("English"),
					),
					Small(
						ID("default-language-hint"),
						Text("Supports all features"),
					),
					Label(
						Input(
							Type("radio"),
							Name("language"),
						),
						Text("French"),
					),
					Label(
						Input(
							Type("radio"),
							Name("language"),
						),
						Text("Mandarin"),
					),
					Label(
						Input(
							Type("radio"),
							Name("language"),
						),
						Text("Thai"),
					),
					Label(
						Input(
							Type("radio"),
							Name("language"),
							Disabled(),
						),
						Text("Dothraki"),
					),
				),
			),
			example(
				"Validation",
				Label(
					Input(
						Type("radio"),
						Name("validation"),
						Aria("invalid", "false"),
						Aria("describedby", "valid-hint"),
					),
					Text("Valid"),
				),
				Small(ID("valid-hint"), Text("Looks good!")),
				Label(
					Input(
						Type("radio"),
						Name("validation"),
						Aria("invalid", "true"),
						Aria("describedby", "invalid-hint"),
						Checked(),
					),
					Text("Invalid"),
				),
				Small(ID("invalid-hint"), Text("This cobination is not allowed!")),
			),
		),
	})
}
