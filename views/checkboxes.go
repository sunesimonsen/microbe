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
						),
						Text("English"),
					),
					Small(Text("You can't disable the default language")),
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
		),
	})
}
