package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ButtonView() Node {
	return docpage(
		HGroup(H1(Text("Button"))),
		NewExample(
			"Styles",
			Nodes(
				Class("grid"),
				Button(Class("solid"), Text("Solid Button")),
				Button(Class("outline"), Text("Outline Button")),
				Button(Class("ghost"), Text("Ghost Button")),
			),
			WithDescription(
				P(
					Text("Buttons are not styled by default to avoid the need to reset styles for custom usages. "),
					InlineCodeList(".solid", ".outline", ".ghost"),
					Text(" to set the appearance of the button."),
				),
			),
		),
		NewExample(
			"Disabled",
			Nodes(
				Class("grid"),
				Button(Class("solid"), Disabled(), Text("Solid Button")),
				Button(Class("outline"), Disabled(), Text("Outline Button")),
				Button(Class("ghost"), Disabled(), Text("Ghost Button")),
			),
		),
		NewExample(
			"Media",
			Nodes(
				Class("grid"),
				Button(Class("outline"), gearIcon(), Text("Settings")),
				Button(Class("solid"), Text("Settings"), gearIcon()),
			),
			WithDescription(
				P(Text("Buttons uses flexbox layout to neatly position icons.")),
			),
		),
		NewExample(
			"Icon",
			Nodes(
				Class("grid small"),
				Button(Class("solid icon"), Aria("label", "Extra options"), burgerIcon()),
				Button(Class("outline icon"), Aria("label", "Extra options"), burgerIcon()),
				Button(Class("ghost icon"), Aria("label", "Extra options"), burgerIcon()),
			),
			WithDescription(
				P(Text("Use the "), Code(Text(".icon")), Text(" for buttons with only an icon.")),
			),
		),
		NewExample(
			"Types",
			Nodes(
				Class("rows"),
				Div(
					Class("grid stretch"),
					Button(Class("solid"), Text("Button")),
					Button(Class("outline"), Text("Button")),
					Button(Class("ghost"), Text("Button")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("button"), Value("Input input")),
					Input(Class("outline"), Type("button"), Value("Input input")),
					Input(Class("ghost"), Type("button"), Value("Input input")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("submit"), Value("Summit input")),
					Input(Class("outline"), Type("submit"), Value("Summit input")),
					Input(Class("ghost"), Type("submit"), Value("Summit input")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("reset"), Value("Reset input")),
					Input(Class("outline"), Type("reset"), Value("Reset input")),
					Input(Class("ghost"), Type("reset"), Value("Reset input")),
				),
			),
			WithDescription(
				P(
					Text("In addition to the "), InlineCodeList("button"), Text(" element, input elements with a type of "),
					InlineCodeList("button", "submit", "reset"),
					Text(" can be styled as a button using the classes "),
					InlineCodeList(".solid", ".outline", ".ghost"),
					Text("."),
				),
			),
		),
	)
}
