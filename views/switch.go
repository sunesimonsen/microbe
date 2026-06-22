package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SwitchView() Node {
	return docpage(
		HGroup(H1(Text("Switch"))),
		NewExample(
			"Example",
			Nodes(
				Label(
					Input(Type("checkbox"), Name("terms"), Role("switch")),
					Text("I agree to the Terms"),
				),
				Label(
					Input(Type("checkbox"), Name("opt-in"), Role("switch"), Checked()),
					Text("Receive news and offers"),
				),
			),
		),
		NewExample(
			"Disabled",
			Nodes(
				Label(
					Input(
						Type("checkbox"),
						Role("switch"),
						Disabled(),
						Checked(),
					),
					Text("Disabled"),
				),
				Label(
					Input(
						Type("checkbox"),
						Role("switch"),
						Disabled(),
					),
					Text("Disabled"),
				),
			),
		),
		NewExample(
			"Hint",
			Nodes(
				Label(
					Input(
						Type("checkbox"),
						Name("newsletter"),
						Role("switch"),
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
		),
		NewExample(
			"Validation",
			Nodes(
				Label(
					Input(
						Type("checkbox"),
						Role("switch"),
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
						Role("switch"),
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
	)
}
